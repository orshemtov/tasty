package main

import (
	"net/http"
	"os"
)

const baseUrl = "https://api.tastyworks.com"

var (
	client        = http.Client{}
	username      = os.Getenv("TASTYWORKS_USERNAME")
	password      = os.Getenv("TASTYWORKS_PASSWORD")
	accountNumber = os.Getenv("TASTYWORKS_ACCOUNT_NUMBER")
	token         string
)

func init() {
	c, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	token = c.Token
}
