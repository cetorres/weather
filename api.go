package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const WEATHER_API_URL = "https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s"

type Weather struct {
	Weather []struct {
		Main string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
		Pressure int64 `json:"pressure"`
		Humidity int64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
	Timezone int `json:"timezone"`
}

func getWeather(city string, unit string) Weather {
	if unit == "C" {
		unit = "metric"
	} else {
		unit = "imperial"
	}

	res, err := http.Get(fmt.Sprintf(WEATHER_API_URL, city, os.Getenv("WEATHER_API_KEY"), unit))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		if res.StatusCode == 401 {
			printError(fmt.Sprintf("Could not access the Weather API (%s).", res.Status))
		} else if res.StatusCode == 404 {
			printError(fmt.Sprintf("Could not find city \"%s\" (%s).", city, res.Status))
		} else {
			printError(res.Status)
		}
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		printError(fmt.Sprintf("Could not read the API response. %s.", err))
		os.Exit(1)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		printError(fmt.Sprintf("Could not convert the API response. %s.", err))
		os.Exit(1)
	}

	return weather
}