package models

import "time"

type Checkout struct {
	CheckoutId       string    `json:"id" firestore:"id"`
	ProductCart      Cart      `json:"cart" firestore:"cart"`
	Payer            User      `json:"payer" firestore:"payer"`
	DeliveryTime     int       `json:"deliveryTime" firestore:"deliveryTime"`
	DeliveryFee      float64   `json:"deliveryFee" firestore:"deliveryFee"`
	Payment          Payment   `json:"payment" firestore:"payment"`
	RemainingSeconds int       `json:"remainingSeconds" firestore:"remainingSeconds"`
	IsTimerStarted   bool      `json:"isTimerStarted" firestore:"isTimerStarted"`
	CreartedAt       time.Time `json:"createdAt" firestore:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt" firestore:"updatedAt"`
}
