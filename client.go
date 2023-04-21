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
	accountNumber string
	token         string
)

func init() {
	acc, err := readAccountNumber()
	if err != nil {
		panic(err)
	}
	accountNumber = acc

	t, err := readToken()
	if err != nil {
		panic(err)
	}
	token = t
}
