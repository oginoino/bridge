package config

import (
	firestore "cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
)

var logger *Logger
var authClient *auth.Client
var dbClient *firestore.Client

func Init() error {
	var err error
	var logger *Logger

	dbClient, err = InitializeDB()

	if err != nil {
		logger.Errorf("failed to initialize Firestore: %v", err)
	}

	authClient, err = InitializeFirebaseAuth()

	if err != nil {
		logger.Errorf("failed to initialize Firebase Auth: %v", err)
	}

	return nil
}

func GetDbClient() *firestore.Client {
	return dbClient
}

func GetAuthClient() *auth.Client {
	return authClient
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
