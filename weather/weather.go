package weather

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather/geo"
	"weather/output"
)

func GetWeather(geo geo.GeoData, format int) string {
	envUrl := os.Getenv("WEATHER_URL")
	baseUrl, err := url.Parse(envUrl + geo.City)

	if err != nil {
		output.PrintError("Некорректный URL сервиса погоды", err)
		return ""
	}

	params := url.Values{}
	params.Add("format", string(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		output.PrintError("Ошибка запроса погоды", err)
		return ""
	}

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		output.PrintError("Некорректный статус ответа при запросе погоды", err)
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output.PrintError("Некорректное тело ответа на запрос погоды", err)
		return ""
	}
	return string(body)

}
