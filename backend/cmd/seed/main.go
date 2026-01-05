package main

import (
	"backend/internal/infra/ent"
	"backend/pkg/config"
	"context"
	"log"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.MustLoadConfig()

	client, err := ent.Open("mysql", cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed opening connection to mysql", "error", err)
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()

	if _, err := client.User.Delete().Exec(ctx); err != nil {
		slog.Error("failed to clear existing users", "error", err)
		log.Fatal(err)
	}
	slog.Info("Cleared existing users")

	// Seed users
	users := []struct {
		Name  string
		Email string
	}{
		{"user1", "user1@example.com"},
		{"user2", "user2@example.com"},
		{"user3", "user3@example.com"},
		{"user4", "user4@example.com"},
		{"user5", "user5@example.com"},
	}

	for _, u := range users {
		_, err := client.User.Create().
			SetName(u.Name).
			SetEmail(u.Email).
			Save(ctx)
		if err != nil {
			slog.Error("failed to create user", "name", u.Name, "error", err)
			continue
		}
		slog.Info("Created user", "name", u.Name, "email", u.Email)
	}

	slog.Info("Seed data inserted successfully")
}
