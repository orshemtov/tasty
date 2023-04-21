package main

import "fmt"

func main() {
	r, err := Accounts()
	if err != nil {
		panic(err)
	}
	for _, a := range r.Data.Items {
		fmt.Printf("account ID: %s\n", a.Account.AccountNumber)
	}
}
