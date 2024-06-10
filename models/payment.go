package models

import "time"

type Payment struct {
	PaymentId     string        `json:"id" firestore:"id"`
	PaymentMethod PaymentMethod `json:"payment_method" firestore:"payment_method"`
	Status        string        `json:"status" firestore:"status"`
	Items         Cart          `json:"items" firestore:"items"`
	Payer         User          `json:"payer" firestore:"payer"`
	Total         float64       `json:"total" firestore:"total"`
	CreatedAt     time.Time     `json:"created_at" firestore:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" firestore:"updated_at"`
}

type PaymentMethod struct {
	MethodType        string      `json:"type" firestore:"type"`
	MethodDescription string      `json:"description" firestore:"description"`
	MethodData        interface{} `json:"data" firestore:"data"`
}
