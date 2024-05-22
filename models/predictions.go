package models

type Predictions struct {
	Predictions  []Prediction `json:"predictions"`
	Status       string       `json:"status"`
	ErrorMessage string       `json:"error_message,omitempty"`
}

type Prediction struct {
	Description string               `json:"description"`
	PlaceID     string               `json:"place_id"`
	Structured  StructuredFormatting `json:"structured_formatting"`
}

type StructuredFormatting struct {
	MainText      string `json:"main_text"`
	SecondaryText string `json:"secondary_text"`
}
