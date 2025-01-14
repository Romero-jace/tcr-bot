package app

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Black-And-White-Club/tcr-bot/app/eventbus"
	"github.com/Black-And-White-Club/tcr-bot/app/modules/leaderboard"
	leaderboardevents "github.com/Black-And-White-Club/tcr-bot/app/modules/leaderboard/domain/events"
	"github.com/Black-And-White-Club/tcr-bot/app/modules/user"
	userevents "github.com/Black-And-White-Club/tcr-bot/app/modules/user/domain/events"
	"github.com/Black-And-White-Club/tcr-bot/app/shared"
	"github.com/Black-And-White-Club/tcr-bot/config"
	"github.com/Black-And-White-Club/tcr-bot/db/bundb"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// App holds the application components.
type App struct {
	Config            *config.Config
	Logger            *slog.Logger
	Router            *message.Router
	UserModule        *user.Module
	LeaderboardModule *leaderboard.Module
	DB                *bundb.DBService
	EventBus          shared.EventBus
}

// Initialize initializes the application.
func (app *App) Initialize(ctx context.Context) error {
	// Parse config file
	configFile := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	app.Config = cfg

	// Initialize logger
	app.Logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app.Logger.Info("App Initialize started")

	// Initialize database
	app.DB, err = bundb.NewBunDBService(ctx, cfg.Postgres)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize EventBus
	app.EventBus, err = eventbus.NewEventBus(ctx, cfg.NATS.URL, app.Logger)
	if err != nil {
		return fmt.Errorf("failed to create event bus: %w", err)
	}

	// Create the Watermill router
	watermillLogger := watermill.NewSlogLogger(app.Logger)
	router, err := message.NewRouter(message.RouterConfig{}, watermillLogger)
	if err != nil {
		return fmt.Errorf("failed to create Watermill router: %w", err)
	}
	app.Router = router

	// Create streams before initializing modules
	if err := app.EventBus.CreateStream(context.Background(), userevents.UserStreamName); err != nil {
		return fmt.Errorf("failed to create user stream: %w", err)
	}
	if err := app.EventBus.CreateStream(context.Background(), leaderboardevents.LeaderboardStreamName); err != nil {
		return fmt.Errorf("failed to create leaderboard stream: %w", err)
	}

	// Initialize User Module
	userModule, err := user.NewUserModule(ctx, cfg, app.Logger, app.DB.UserDB, app.EventBus)
	if err != nil {
		return fmt.Errorf("failed to initialize user module: %w", err)
	}
	app.UserModule = userModule

	// Initialize Leaderboard Module
	leaderboardModule, err := leaderboard.NewLeaderboardModule(ctx, cfg, app.Logger, app.DB.LeaderboardDB, app.EventBus)
	if err != nil {
		return fmt.Errorf("failed to initialize leaderboard module: %w", err)
	}
	app.LeaderboardModule = leaderboardModule

	// Wait for subscribers to be ready
	<-userModule.SubscribersReady
	<-leaderboardModule.SubscribersReady

	// Add a delay here
	time.Sleep(100 * time.Millisecond)

	app.Logger.Info("All modules initialized successfully")

	return nil
}

// Run starts the application.
func (app *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup

	// Handle OS signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		app.Logger.Info("Interrupt signal received, shutting down...")
		cancel()
	}()

	// Start the router
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.Router.Run(ctx); err != nil {
			app.Logger.Error("Error running Watermill router", slog.Any("error", err))
			cancel()
		}
	}()

	// Wait for graceful shutdown
	wg.Wait()

	app.Close()
	app.Logger.Info("Graceful shutdown complete.")

	return nil
}

// Close shuts down all resources.
func (app *App) Close() {
	if app.Router != nil {
		if err := app.Router.Close(); err != nil {
			app.Logger.Error("Error closing Watermill router", slog.Any("error", err))
		}
	}

	if app.EventBus != nil {
		if err := app.EventBus.Close(); err != nil {
			app.Logger.Error("Error closing event bus", slog.Any("error", err))
		}
	}
}
