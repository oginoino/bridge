package models

import (
	"time"

	"cloud.google.com/go/firestore"
)

type User struct {
	Id              firestore.DocumentRef    `json:"id" firestore:"_id"`
	Uid             string                   `json:"uid" firestore:"uid" validate:"required"`
	UserDisplayName string                   `json:"userDisplayName" firestore:"userDisplayName" validate:"required"`
	UserEmail       string                   `json:"userEmail" firestore:"userEmail" validate:"required" unique:"true"`
	UserPhotoUrl    string                   `json:"userPhotoUrl" firestore:"userPhotoUrl"`
	UserName        string                   `json:"userName" firestore:"userName"`
	CreatedAt       time.Time                `json:"createdAt" firestore:"createdAt"`
	UpdatedAt       time.Time                `json:"updatedAt" firestore:"updatedAt"`
	DeletedAt       time.Time                `json:"deletedAt" firestore:"deletedAt"`
	UserProperties  []map[string]interface{} `json:"userProperties" firestore:"userProperties"`
	IsActivated     bool                     `json:"isActivated" firestore:"isActivated"`
	Addresses       []Address                `json:"addresses" firestore:"addresses"`
	SelectedAddress Address                  `json:"selectedAddress" firestore:"selectedAddress"`
}
