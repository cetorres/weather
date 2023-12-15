package main

import (
	"strings"
	"time"
)

func getIcon(condition string) string {
	icon := ""
	if strings.Contains(strings.ToLower(condition), "snow") {
		icon = "❄️  "
	} else if strings.Contains(strings.ToLower(condition), "few clouds") {
		icon = "🌥️  "
	} else if strings.Contains(strings.ToLower(condition), "clear sky") {
		currentHour := time.Now().Hour()
		if currentHour < 6 || currentHour > 18 {
			icon = "⚪️ "
		} else {
			icon = "☀️ "
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