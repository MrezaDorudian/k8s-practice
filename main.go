package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"os"
)

type config struct {
	Port       string `json:"port"`
	RemoteHost string `json:"remote_host"`
}

type remote struct {
	Host string `json:"host"`
}

type weather struct {
	Current struct {
		Hostname            string   `json:"hostname"`
		Temperature         int      `json:"temperature"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		WindSpeed           int      `json:"wind_speed"`
		Humidity            int      `json:"humidity"`
		Feelslike           int      `json:"feelslike"`
	} `json:"current"`
}

func (c *config) setConfig() {
	jsonFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
	    c.Port = "8080"
		c.RemoteHost = "http://api.weatherstack.com/current?access_key=9a002e43db71f9388084af35b8e4a947&query="
	} else {
        err = json.Unmarshal(jsonFile, &c)
	    if err != nil {
		    panic(err)
	    }
	}
}

func (r *remote) getWeather(c echo.Context) error {
	url := r.Host + c.Param("city")
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	weatherInfo := weather{}
	user, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	weatherInfo.Current.Hostname = user + "@" + hostname

	err = json.Unmarshal(body, &weatherInfo)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, weatherInfo.Current)
}

func (r *remote) setRemote(host string) {
	r.Host = host
}

func main() {
	e := echo.New()
	conf := config{}
	conf.setConfig()
	rem := remote{}
	rem.setRemote(conf.RemoteHost)
	e.GET("/:city", rem.getWeather)
	e.Logger.Fatal(e.Start(":" + conf.Port))
}
