package handler

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/GinoCodeSpace/bridge/config"
	"gopkg.in/go-playground/validator.v9"
)

var dbClient *firestore.Client
var authClient *auth.Client

func InitializeHandler() (*firestore.Client, *auth.Client) {
	dbClient = config.GetDbClient()
	authClient = config.GetAuthClient()
	return dbClient, authClient
}

type DefaultHandler struct {
	collection *firestore.CollectionRef
	ctx        context.Context
	validate   *validator.Validate
}

func NewDefaultHandler(collection *firestore.CollectionRef, ctx context.Context, validate *validator.Validate) *DefaultHandler {
	return &DefaultHandler{
		collection: collection,
		ctx:        ctx,
		validate:   validate,
	}
}

type AuthHandler struct {
	collection *firestore.CollectionRef
	authClient *auth.Client
	ctx        context.Context
}

func NewAuthHandler(collection *firestore.CollectionRef, authClient *auth.Client, ctx context.Context) *AuthHandler {
	return &AuthHandler{
		collection: collection,
		authClient: authClient,
		ctx:        ctx,
	}
}
