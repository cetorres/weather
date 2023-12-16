package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func getIcon(condition string, timezone int) string {
	icon := ""
	if strings.Contains(strings.ToLower(condition), "snow") {
		icon = "❄️  "
	} else if strings.Contains(strings.ToLower(condition), "few clouds") {
		icon = "🌥️  "
	} else if strings.Contains(strings.ToLower(condition), "clear sky") {
		currentHour := time.Now().UTC().Add(time.Second * time.Duration(timezone)).Hour()
		if currentHour < 6 || currentHour >= 18 {
			icon = "⚪️ "
		} else {
			icon = "☀️  "
		}
	} else if strings.Contains(strings.ToLower(condition), "clouds") {
		icon = "☁️  "
	} else if strings.Contains(strings.ToLower(condition), "shower") {
		icon = "🌧️  "
	} else if strings.Contains(strings.ToLower(condition), "rain") {
		icon = "🌦️  "
	} else if strings.Contains(strings.ToLower(condition), "thunderstorm") {
		icon = "⛈️  "
	} else if strings.Contains(strings.ToLower(condition), "drizzle") {
		icon = "🌨️  "
	} else {
		icon = "☁️  "
	}

	return icon
}

func printError(err string) {
	fmt.Print(LOGO)
	fmt.Println(TITLE)
	color.Red(fmt.Sprintf("ERROR: %s", err))
}

func printHelp() {
	fmt.Println(DESCRIPTION)
	if os.Getenv("WEATHER_API_KEY") == "" {
		color.Red("\nERROR: WEATHER_API_KEY is missing.")
		os.Exit(1)
	}
	os.Exit(0)
}