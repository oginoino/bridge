package models

type Product struct {
	Id                      string   `json:"id" firestore:"id" unique:"true" validate:"required"`
	ProductName             string   `json:"productName" firestore:"productName" validate:"required"`
	ProductPrice            float64  `json:"productPrice" firestore:"productPrice" validate:"required"`
	ProductUnitOfMeasure    string   `json:"productUnitOfMeasure" firestore:"productUnitOfMeasure" validate:"required"`
	ProductUnitQuantity     string   `json:"productUnitQuantity" firestore:"productUnitQuantity" validate:"required"`
	ProductCategories       []string `json:"productCategories" firestore:"productCategories" validate:"required"`
	AvailableQuantity       int      `json:"availableQuantity" firestore:"availableQuantity" validate:"required"`
	ContentValue            string   `json:"contentValue" firestore:"contentValue"`
	ProductImageSrc         string   `json:"productImageSrc" firestore:"productImageSrc"`
	ProducKilogramsWeight   float64  `json:"producKilogramsWeight" firestore:"producKilogramsWeight"`
	ProductCubicMeterVolume float64  `json:"productCubicMeterVolume" firestore:"productCubicMeterVolume"`
}
