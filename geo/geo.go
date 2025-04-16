package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"weather/output"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {

	geoData := &GeoData{}

	if city != "" && CheckCity(city) {
		geoData.City = city
		return geoData, nil
	}

	output.PrintInfo("Город не задан и будет определен по IP")
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
	defer resp.Body.Close()

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

	err = json.Unmarshal(body, geoData)
	if err != nil {
		output.PrintError("Не удалось найти город в ответе с геолокацией", err)
		return nil, err
	}

	return geoData, nil
}

type cityPopulationResponse struct {
	Err bool `json:"error"`
}

func CheckCity(city string) bool {
	postBody, err := json.Marshal(map[string]string{
		"city": city,
	})

	if err != nil {
		output.PrintWarning("Не удалось проверить заданный город", err)
		return false
	}

	envUrl := os.Getenv("CHECK_CITY_URL")
	if err != nil {
		output.PrintWarning("Некорректный URL сервиса проверки города", err)
		return false
	}

	resp, err := http.Post(envUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		output.PrintWarning("Ошибка запроса при проверке города", err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output.PrintWarning("Ошибка ответа при проверке города", err)
		return false
	}

	isCity := cityPopulationResponse{}
	err = json.Unmarshal(body, &isCity)
	if err != nil {
		output.PrintWarning("Неудалось разобрать ответ при проверки города", err)
		return false
	}

	if isCity.Err {
		output.PrintInfo("При проверке не удалось подтвердить существование указанного города")
		return false
	}
	return true
}
