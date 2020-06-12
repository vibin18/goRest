package localApiGather

import (
	"encoding/json"
	"log"
	"net/http"
)

type CurrentWeather struct {
	Humidity   string  `json:"Humidity"`
	Preassure  string `json:"Preassure"`
	Temprature string `json:"Temprature"`
	}

var (
	weatherBaseUrl = "http://192.168.178.68:8081"
	weatherSubPath = "/temp"
)

func localApiGather() CurrentWeather {

	//resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?id=2947444&units=metric&appid=3a20a64aa7cb7e38a83c7cb2b48ab460")
	resp, err := http.Get(weatherBaseUrl + weatherSubPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var gatheredResult CurrentWeather
	json.NewDecoder(resp.Body).Decode(&gatheredResult)


	return gatheredResult
}


