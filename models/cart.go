package models

type Cart struct {
	ID           string        `json:"id" firestore:"id"`
	ProductItems []ProductItem `json:"productsItems" firestore:"productsItems"`
	Shopper      User          `json:"shopper" firestore:"shopper"`
}

type ProductItem struct {
	ProductId        string  `json:"productId" firestore:"productId"`
	Product          Product `json:"product" firestore:"product"`
	SelectedQuantity int     `json:"selectedQuantity" firestore:"selectedQuantity"`
}
