package config

import (
	"context"

	firestore "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func InitializeDB() (*firestore.Client, error) {
	logger := GetLogger("FIREBASE FIRESTORE")
	logger.Info("Initializing Firebase Firestore")

	opt := option.WithCredentialsFile("credentials.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	client, err := app.Firestore(context.Background())

	if err != nil {
		logger.Error(err)
		defer client.Close()
		return nil, err
	}

	logger.Info("Initialized Firebase Firestore")

	return client, nil

}
