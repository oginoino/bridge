package models

type Address struct {
	Id            string `json:"id" firestore:"id"`
	Description   string `json:"description" firestore:"description"`
	MainText      string `json:"mainText" firestore:"mainText"`
	SecondaryText string `json:"secondaryText" firestore:"secondaryText"`
}
