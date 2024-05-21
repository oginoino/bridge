package models

type User struct {
	Id              string                   `json:"id" firestore:"id" unique:"true"`
	Uid             string                   `json:"uid" firestore:"uid" validate:"required" unique:"true"`
	UserDisplayName string                   `json:"userDisplayName" firestore:"userDisplayName" validate:"required"`
	UserEmail       string                   `json:"userEmail" firestore:"userEmail" validate:"required" unique:"true"`
	UserPhotoUrl    string                   `json:"userPhotoUrl" firestore:"userPhotoUrl"`
	UserName        string                   `json:"userName" firestore:"userName"`
	CreatedAt       CustomTime               `json:"createdAt" firestore:"createdAt"`
	UpdatedAt       CustomTime               `json:"updatedAt" firestore:"updatedAt"`
	UserProperties  []map[string]interface{} `json:"userProperties" firestore:"userProperties"`
	IsActivated     bool                     `json:"isActivated" firestore:"isActivated"`
	Addresses       []Address                `json:"addresses" firestore:"addresses"`
	SelectedAddress Address                  `json:"selectedAddress" firestore:"selectedAddress"`
}
