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


	resp, err := http.Get(weatherBaseUrl + weatherSubPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var gatheredResult CurrentWeather
	json.NewDecoder(resp.Body).Decode(&gatheredResult)


	return gatheredResult
}


