package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
)

type MapsService struct {
	Language     string
	Key          string
	Location     string
	Radius       string
	StrictBounds bool
}

func NewMapsService() *MapsService {
	return &MapsService{
		Language:     "pt",
		Key:          os.Getenv("M_API_KEY"),
		Location:     "-23.7213129,-46.7565639",
		Radius:       "1000",
		StrictBounds: true,
	}
}

func (ms *MapsService) FetchAddress(input string) (models.Predictions, error) {
	pathPlaceAutoComplete := "https://maps.googleapis.com/maps/api/place/autocomplete/json"

	queryParameters := map[string]string{
		"language":     ms.Language,
		"input":        input,
		"key":          ms.Key,
		"location":     ms.Location,
		"radius":       ms.Radius,
		"strictbounds": "true",
	}

	if !ms.StrictBounds {
		queryParameters["strictbounds"] = "false"
	}

	uri, err := ms.buildURL(pathPlaceAutoComplete, queryParameters)
	if err != nil {
		return models.Predictions{}, err
	}

	client := http.Client{Timeout: 10 * time.Second}
	response, err := client.Get(uri)
	if err != nil {
		return models.Predictions{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ms.handleErrors(response)
	}

	return ms.handleSuccess(response)
}

func (ms *MapsService) buildURL(baseURL string, queryParameters map[string]string) (string, error) {
	uri, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	query := uri.Query()
	for key, value := range queryParameters {
		query.Set(key, value)
	}
	uri.RawQuery = query.Encode()
	return uri.String(), nil
}

func (ms *MapsService) handleSuccess(response *http.Response) (models.Predictions, error) {
	var predictions models.Predictions
	if err := json.NewDecoder(response.Body).Decode(&predictions); err != nil {
		return predictions, err
	}
	return predictions, nil
}

func (ms *MapsService) handleErrors(response *http.Response) (models.Predictions, error) {
	var predictions models.Predictions
	if err := json.NewDecoder(response.Body).Decode(&predictions); err != nil {
		return predictions, err
	}
	if predictions.Status != "OK" {
		return predictions, fmt.Errorf(predictions.ErrorMessage)
	}
	return predictions, fmt.Errorf("unknown error")
}
