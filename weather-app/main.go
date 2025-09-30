package main

import (
	"fmt"
	"golang-study/handler"
)

func main() {

	println("_______________________\n")
	println("👋 Welcome to weather app! 👋")
	println("_______________________\n")

	var city string

	fmt.Print("➡️ Enter a city: ")
	fmt.Scan(&city)
	println("_______________________\n")
	
	weatherData := handler.GetWeather(city)

	
	fmt.Printf("Weather at city %v - 🌤️ %v\n", city, weatherData.Temperature)
	fmt.Println(weatherData.Description)
	
}