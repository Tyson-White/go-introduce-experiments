package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ForeCast struct {
	Day string `json:"day"`
	Temperature string `json:"temperature"`
	Wind string `json:"wind"`
}

type WeatherData struct {
	Temperature string `json:"temperature"`
	Wind string `json:"wind"`
	Description string `json:"description"`
	Forecast []ForeCast `json:"forecast"`
}

func GetWeather(city string) WeatherData  {

	fmt.Println("🔍 Searching...")
	println("_______________________\n")
	resp, err := http.Get(fmt.Sprintf("https://goweather.herokuapp.com/weather/%v", city))

	if err != nil {

	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var data WeatherData

	if err := json.Unmarshal(body, &data); err != nil {
    }


	return data
}