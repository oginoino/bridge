package models

import (
	"cloud.google.com/go/firestore"
)

type Address struct {
	Id            firestore.DocumentRef `json:"id" firestore:"_id"`
	Description   string                `json:"description" firestore:"description"`
	MainText      string                `json:"mainText" firestore:"mainText"`
	SecondaryText string                `json:"secondaryText" firestore:"secondaryText"`
}
