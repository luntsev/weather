package geo_test

import (
	"os"
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	os.Setenv("CHECK_CITY_URL", "https://countriesnow.space/api/v0.1/countries/population/cities")
	city := "london"
	want := geo.GeoData{City: city}

	got, err := geo.GetMyLocation(city)
	if err != nil {
		t.Error("Ошибка при выполнении функции GetMyLocation")
	}

	if got.City != want.City {
		t.Errorf("Ожидалось: %v, вернулось: %v", want, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	os.Setenv("CHECK_CITY_URL", "https://countriesnow.space/api/v0.1/countries/population/cities")
	city := "moscowpiter"
	want := geo.ErrNoCity
	_, got := geo.GetMyLocation(city)
	if got != want {
		t.Errorf("Ожидалось: %s, вернулось: %s", want, got)
	}
}

func TestGetMyLocationBadGeoURL(t *testing.T) {
	os.Setenv("GEO_URL", "htp\\:/countriesnow?")
	city := ""
	want := geo.ErrBadGeoUrl
	_, got := geo.GetMyLocation(city)
	if got != want {
		t.Errorf("Ожидалось: %s, вернулось: %s", want, got)
	}
}
