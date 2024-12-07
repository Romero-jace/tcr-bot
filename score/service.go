package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Black-And-White-Club/tcr-bot/app/models"
	"github.com/Black-And-White-Club/tcr-bot/internal/db"
	"github.com/Black-And-White-Club/tcr-bot/internal/nats"
)

// ScoreService handles score-related database operations.
type ScoreService struct {
	db                 db.ScoreDB
	natsConnectionPool *nats.NatsConnectionPool
}

// NewScoreService creates a new ScoreService.
func NewScoreService(db db.ScoreDB, natsConnectionPool *nats.NatsConnectionPool) *ScoreService {
	return &ScoreService{
		db:                 db,
		natsConnectionPool: natsConnectionPool,
	}
}

// GetUserScore retrieves the score for a specific user and round.
func (s *ScoreService) GetUserScore(ctx context.Context, discordID, roundID string) (*models.Score, error) {
	return s.db.GetUserScore(ctx, discordID, roundID)
}

// GetScoresForRound retrieves all scores for a given round.
func (s *ScoreService) GetScoresForRound(ctx context.Context, roundID string) ([]models.Score, error) {
	return s.db.GetScoresForRound(ctx, roundID)
}

// ProcessScores processes a batch of scores, preparing them for insertion and publishing an event.
func (s *ScoreService) ProcessScores(ctx context.Context, roundID int64, scores []models.ScoreInput) error {
	// Sort the scores in descending order
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].Score > scores[j].Score
	})

	// Assign tag numbers based on ranking
	for i := range scores {
		tagNumber := i + 1
		scores[i].TagNumber = &tagNumber
	}

	modelScores := make([]models.Score, len(scores))
	for i, score := range scores {
		modelScores[i] = models.Score{
			DiscordID: score.DiscordID,
			RoundID:   strconv.FormatInt(roundID, 10),
			Score:     score.Score,
			TagNumber: *score.TagNumber,
		}
	}

	err := s.db.ProcessScores(ctx, roundID, modelScores) // Simplified database insertion
	if err != nil {
		return err
	}
	// Publish ScoresProcessedEvent to the leaderboard module
	err = s.natsConnectionPool.Publish("scores.processed", &nats.ScoresProcessedEvent{
		RoundID:         roundID,
		ProcessedScores: scores,
	})
	if err != nil {
		return fmt.Errorf("failed to publish scores.processed event: %w", err)
	}

	return nil
}

// UpdateScore updates a specific score.
func (s *ScoreService) UpdateScore(ctx context.Context, roundID, discordID string, score int, tagNumber *int) (*models.Score, error) {
	existingScore, err := s.db.GetUserScore(ctx, discordID, roundID)
	if err != nil {
		return nil, err
	}
	if existingScore == nil {
		return nil, errors.New("score not found")
	}

	existingScore.Score = score
	if tagNumber != nil {
		existingScore.TagNumber = *tagNumber
	}

	err = s.db.UpdateScore(ctx, existingScore) // Simplified database update
	if err != nil {
		return nil, fmt.Errorf("failed to update score: %w", err)
	}

	return existingScore, nil
}

// StartNATSSubscribers starts the NATS subscribers for the score service.
func (s *ScoreService) StartNATSSubscribers(ctx context.Context) error {
	conn, err := s.natsConnectionPool.GetConnection()
	if err != nil {
		return fmt.Errorf("failed to get NATS connection from pool: %w", err)
	}
	defer s.natsConnectionPool.ReleaseConnection(conn)

	// Subscribe to "round.finalized" subject
	_, err = conn.Subscribe("round.finalized", func(msg *nats.Msg) {
		var event nats.RoundFinalizedEvent
		err := json.Unmarshal(msg.Data, &event)
		if err != nil {
			log.Printf("Error unmarshaling RoundFinalizedEvent: %v", err)
			return
		}

		// Retrieve the scores for the finalized round
		roundID := event.RoundID

		// Get all scores for the round
		scores, err := s.db.GetScoresForRound(ctx, strconv.FormatInt(roundID, 10))
		if err != nil {
			log.Printf("Error getting scores for round: %v", err)
			return
		}

		// Convert the scores to ScoreInput
		var scoreInputs []models.ScoreInput
		for _, score := range scores {
			scoreInputs = append(scoreInputs, models.ScoreInput{
				DiscordID: score.DiscordID,
				Score:     score.Score,
			})
		}

		// Process the scores
		err = s.ProcessScores(ctx, roundID, scoreInputs)
		if err != nil {
			log.Printf("Error processing scores: %v", err)
			return
		}
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to round.finalized: %w", err)
	}

	return nil
}