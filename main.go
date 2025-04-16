package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/output"
	"weather/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалосб загрузить переменные окружения", err)
	}
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		output.PrintError("Не удалось определить геолокацию", err)
	}

	weatherData := weather.GetWeather(*geoData, *format)

	fmt.Println(weatherData)
}
