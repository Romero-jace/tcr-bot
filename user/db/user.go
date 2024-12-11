// user/db/user.go
package userdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/Black-And-White-Club/tcr-bot/models"
	"github.com/uptrace/bun"
)

// userDBImpl is an implementation of the UserDB interface using bun.
type UserDBImpl struct {
	DB *bun.DB
}

// CreateUser creates a new user.
func (db *UserDBImpl) CreateUser(ctx context.Context, user *models.User) error { // Use *models.User
	if db.DB == nil {
		return errors.New("database connection is not initialized")
	}

	if user == nil {
		return errors.New("user cannot be nil")
	}

	_, err := db.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUser retrieves a user by Discord ID.
func (db *UserDBImpl) GetUserByDiscordID(ctx context.Context, discordID string) (*User, error) {
	var user User
	err := db.DB.NewSelect().
		Model(&user).
		Where("discord_id = ?", discordID).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// UpdateUser updates an existing user.
func (db *UserDBImpl) UpdateUser(ctx context.Context, discordID string, updates *User) error {
	_, err := db.DB.NewUpdate().
		Model(updates).
		Column("name", "role").
		Where("discord_id = ?", discordID).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
