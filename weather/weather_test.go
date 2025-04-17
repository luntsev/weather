package weather_test

import (
	"os"
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

var cityTestCases = []struct {
	city string
}{
	{city: "Moscow"},
	{city: "London"},
	{city: "Berlin"},
	{city: "Madrid"},
}

func TestGetWeatherCurrentCity(t *testing.T) {
	os.Setenv("CHECK_CITY_URL", "https://countriesnow.space/api/v0.1/countries/population/cities")
	os.Setenv("WEATHER_URL", "https://wttr.in/")
	os.Setenv("GEO_URL", "https://wttr.in/")
	for _, tc := range cityTestCases {
		t.Run(tc.city, func(t *testing.T) {
			want := tc.city
			geo := geo.GeoData{City: want}
			result, err := weather.GetWeather(geo, 3)
			if err != nil {
				t.Error(err)
			}
			got := strings.TrimSuffix(strings.Fields(result)[0], ":")
			if got != want {
				t.Errorf("Ожидалось: %s, вернулось: %s", want, got)
			}
		})
	}
}

var formatTestCases = []struct {
	name   string
	format int
}{
	{name: "Short weather", format: 1},
	{name: "Full weather", format: 2},
	{name: "Short weather with city", format: 3},
	{name: "Full weather with city", format: 4},
}

func TestGetWeatherFormat(t testing.T) {
	os.Setenv("CHECK_CITY_URL", "https://countriesnow.space/api/v0.1/countries/population/cities")
	os.Setenv("WEATHER_URL", "https://wttr.in/")
	os.Setenv("GEO_URL", "https://wttr.in/")
	for _, tc := range formatTestCases {
		t.Run(tc.name, func(t *testing.T) {

			want := weather.ErrBadFormat
			geo := geo.GeoData{City: "Moscow"}
			_, gotErr := weather.GetWeather(geo, tc.format)
			if want != gotErr {
				t.Errorf("Ожидалось: %s, вернулось: %s", want, gotErr)
			}
		})
	}
}

var wrongFormatTestCases = []struct {
	name   string
	format int
}{
	{name: "More that 4", format: 123},
	{name: "Less than 0", format: -321},
	{name: "Equal to 0", format: 0},
}

func TestGetWeatherBadFormat(t *testing.T) {
	os.Setenv("CHECK_CITY_URL", "https://countriesnow.space/api/v0.1/countries/population/cities")
	os.Setenv("WEATHER_URL", "https://wttr.in/")
	os.Setenv("GEO_URL", "https://wttr.in/")
	for _, tc := range wrongFormatTestCases {
		t.Run(tc.name, func(t *testing.T) {
			want := weather.ErrBadFormat
			geo := geo.GeoData{City: "Moscow"}
			_, gotErr := weather.GetWeather(geo, tc.format)
			if want != gotErr {
				t.Errorf("Ожидалось: %s, вернулось: %s", want, gotErr)
			}
		})
	}
}
