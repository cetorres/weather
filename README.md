# weather
A simple terminal command made in Go to get the current weather conditions.

## Build and usage
To build and run the program, just run:

```sh
$ go build
$ ./weather
```

## Weather API
I used the [OpenWeather API](https://openweathermap.org) to query the weather conditions. You will need to obtain your own API key and set an environment variable with it:

```sh
$ export WEATHER_API_KEY=<your_key_here>
```

Get your free API key at https://home.openweathermap.org/users/sign_up. Create your account, then go to API keys.

## Output

```
┬ ┬┌─┐┌─┐┌┬┐┬ ┬┌─┐┬─┐
│││├┤ ├─┤ │ ├─┤├┤ ├┬┘
└┴┘└─┘┴ ┴ ┴ ┴ ┴└─┘┴└─
weather 1.0 - Copyright © 2023 Carlos E. Torres (https://github.com/cetorres)
A simple terminal command to get the current weather conditions.

Usage:
        weather <city> [unit]

Units:
        C: Celcius
        F: Fahrenheit (default)

Pro tip:
        You can set environment variables:
        - export WEATHER_CITY_NAME="New York"
        - export WEATHER_UNIT=C
        And avoid typing parameters everytime.

Uses the OpenWeather API (https://openweathermap.org):
        - Get your API key at https://home.openweathermap.org/users/sign_up
        - Set an environment variable: export WEATHER_API_KEY=<your_key_here>
```

### Run with a city

![Screenshot](screenshot.png)
