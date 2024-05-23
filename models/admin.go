package models

type AdminUser struct {
	User
	Role string `json:"role" firestore:"role"`
}
