package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Time struct {
	Updated    string
	UpdatedISO string
	Updateduk  string
}
type Cury struct {
	Code       string
	Symbol     string
	Rate       string
	Rate_Float float64
}
type Bpi struct {
	USD Cury
	EUR Cury
	GBP Cury
}
type ApiRes struct {
	Time      Time
	ChartName string
	Bpi       Bpi
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		response, _ := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
		data, _ := ioutil.ReadAll(response.Body)
		var res ApiRes

		json.Unmarshal(data, &res)
		veri, _ := json.Marshal(res)

		return c.Format(veri)
	})
	app.Listen(":8080")
}
