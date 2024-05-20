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

	opt := option.WithCredentialsFile("credentials.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Errorf("Failed to initialize Firebase Auth: %v", err)
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		logger.Errorf("Failed to initialize Firebase Auth: %v", err)
		return nil, err
	}

	logger.Info("Initialized Firebase Auth")

	return client, nil
}
