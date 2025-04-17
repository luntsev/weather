package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather/geo"
	"weather/output"
)

var ErrBadWthURL = errors.New("bad weather service url")
var ErrBadFormat = errors.New("bad format")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	envUrl := os.Getenv("WEATHER_URL")

	if format < 1 || format > 4 {
		return "", ErrBadFormat
	}

	baseUrl, err := url.Parse(envUrl + geo.City)
	if err != nil {
		output.PrintError("Некорректный URL сервиса погоды", err)
		return "", ErrBadWthURL
	}

	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	fmt.Println(baseUrl.String())

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		output.PrintError("Ошибка запроса погоды", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err := errors.New(resp.Status)
		output.PrintError("Некорректный статус ответа при запросе погоды", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output.PrintError("Некорректное тело ответа на запрос погоды", err)
		return "", err
	}
	return string(body), nil
}
