/*
Weather
A simple terminal command to get the current weather conditions.
Copyright © 2023 Carlos E. Torres <cetorres@cetorres.com>
https://github.com/cetorres
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

const LOGO = `┬ ┬┌─┐┌─┐┌┬┐┬ ┬┌─┐┬─┐
│││├┤ ├─┤ │ ├─┤├┤ ├┬┘
└┴┘└─┘┴ ┴ ┴ ┴ ┴└─┘┴└─
`
const TITLE = `weather 1.0 - Copyright © 2023 Carlos E. Torres (https://github.com/cetorres)`
var DESCRIPTION = fmt.Sprintf(`%s%s
A simple terminal command to get the current weather conditions.

Usage:
	weather <city> [unit]

Units:
	C: Celcius
	F: Fahrenheit (default)

Pro tip:
	You can set environment variables:
	- WEATHER_CITY_NAME="New York"
	- WEATHER_UNIT=C
	And avoid typing parameters everytime.

Uses the OpenWeather API (https://openweathermap.org)`, LOGO, TITLE)

func main() {
	// Check for WEATHER_API_KEY
	if WEATHER_API_KEY == "" {
		fmt.Printf("%s%s\n\n", LOGO, TITLE)
		color.Red("WEATHER_API_KEY is missing.")
		color.Yellow("-> Get yours at https://home.openweathermap.org/users/sign_up\n\n")
		os.Exit(0)
	}

	// Check arguments and show help
	if len(os.Args) < 2 && os.Getenv("WEATHER_CITY_NAME") == "" {
		fmt.Println(DESCRIPTION)
		os.Exit(0)
	}
	
	// Get city name
	city := ""
	if os.Getenv("WEATHER_CITY_NAME") != "" {
		city = os.Getenv("WEATHER_CITY_NAME")
	} else {
		city = os.Args[1]
	}
	
	// Get unit
	unit := "F"
	speedUnit := "mph"
	if os.Getenv("WEATHER_UNIT") != "" {
		unit = strings.ToUpper(os.Getenv("WEATHER_UNIT"))
	} else {
		if len(os.Args) == 3 {
			unit = strings.ToUpper(os.Args[2])
		}
	}
	if unit == "C" {
		speedUnit = "m/s"
	}
	
	// Get weather conditions
	weather := getWeather(city, unit)

	// Get condition icon
	icon := getIcon(weather.Weather[0].Description)

	// Print weather conditions
	titleColor := color.New(color.BgHiBlue, color.Bold)
	titleColor.Println(fmt.Sprintf(" %s, %s | %.0fᵒ%s %s%s ", city, weather.Sys.Country, weather.Main.Temp, strings.ToUpper(unit), icon, weather.Weather[0].Description))
	color.HiGreen(fmt.Sprintf("⚬ Feels like: %.0fᵒ%s", weather.Main.FeelsLike, unit))
	color.HiBlue(fmt.Sprintf("⚬ Min Temp: %.0fᵒ%s", weather.Main.TempMin, unit))
	color.Red(fmt.Sprintf("⚬ Max Temp: %.0fᵒ%s", weather.Main.TempMax, unit))
	color.White(fmt.Sprintf("⚬ Wind speed: %.0f %s", weather.Wind.Speed, speedUnit))
	color.White(fmt.Sprintf("⚬ Pressure: %d hPa", weather.Main.Pressure))
	color.White(fmt.Sprintf("⚬ Humidity: %d%%", weather.Main.Humidity))
}
