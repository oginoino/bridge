package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitializeFirebaseAuth() (*auth.Client, error) {
	logger := GetLogger("FIREBASE AUTH")
	logger.Info("Initializing Firebase Auth")

	ctx := context.Background()

	opt := option.WithCredentialsFile("credentials.json")

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logger.Errorf("Failed to initialize Firebase App: %v", err)
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		logger.Errorf("Failed to initialize Firebase Auth client: %v", err)
		return nil, err
	}

	logger.Info("Initialized Firebase Auth successfully")

	return client, nil
}
