package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather/output"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		output.PrintInfo("Город не задан и будет определен по IP")
		return &GeoData{City: city}, nil
	}

	geoUrl := os.Getenv("GEO_URL")
	_, err := url.Parse(geoUrl)
	if err != nil {
		output.PrintError("URL сервиса определения геолокации по IP в переменных окружения некорректен", err)
	}

	resp, err := http.Get(geoUrl)
	if err != nil {
		output.PrintError("Ошибка запроса определения геолокации по IP", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		output.PrintError("Некорректный статус ответа при определении геолокации по IP", err)
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output.PrintError("Некорректное тело ответа при определении геолокации по IP", err)
		return nil, err
	}

	geoData := &GeoData{}

	err = json.Unmarshal(body, geoData)
	if err != nil {
		output.PrintError("Не удалось найти город в ответе с геолокацией", err)
		return nil, err
	}

	return geoData, nil
}
