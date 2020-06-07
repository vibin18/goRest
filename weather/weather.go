package weather

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
)

type CurrentWeather struct {
	Temp      float64
	Feels_Like float64
	Temp_Min   float64
	Temp_Max   float64
	Pressure  float64
	Humidity  float64

}

//"main":{"temp":13.75,"feels_like":12.99,"temp_min":13.33,"temp_max":14,"pressure":1011,"humidity":76}

type CurrentLocation struct {
	Type    int
	ID      int
	Country string
	City string
	Sunrise int
	Sunset  int
}

type WeatherLocation interface{}

type results  map[string]interface{}

var (
	weatherBaseUrl = "https://api.openweathermap.org"
	weatherSubPath = "/data/2.5/weather"
	defaultLocationId = "2947444"
	defaultUnitType =  "metric"
	defaultToken = "3a20a64aa7cb7e38a83c7cb2b48ab460"
)


func weatherApiGather( unitType, locationId, token string ) results {

	//resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?id=2947444&units=metric&appid=3a20a64aa7cb7e38a83c7cb2b48ab460")
	resp, err := http.Get(weatherBaseUrl + weatherSubPath + "?id=" + locationId + "&units=" + unitType + "&appid=" + token)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	var gatheredResult results
	json.NewDecoder(resp.Body).Decode(&gatheredResult)

	return gatheredResult
}

func populateData(inData results) (lStruct CurrentLocation, wStruct CurrentWeather )  {

	location := inData["sys"].(map[string]interface{})
	location["city"] = inData["name"]
	weather := inData["main"].(map[string]interface{})

	var locationStruct CurrentLocation
	var weatherStruct CurrentWeather
	mapstructure.Decode(location, &locationStruct)
	mapstructure.Decode(weather, &weatherStruct)
	return locationStruct, weatherStruct
}


func GetWeatherAndLocation() WeatherLocation {

	gatheredData := weatherApiGather(defaultUnitType, defaultLocationId, defaultToken)

	location, weather := populateData(gatheredData)

	combinedWeatherData := struct {
		CurrentWeather
		CurrentLocation
	}{weather,location}

	return combinedWeatherData

}


