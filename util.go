package main

import (
	"encoding/json"
	"os"
)

func SaveToken() error {
	r, err := Auth()
	if err != nil {
		return err
	}

	f, err := os.Create(".cache/token.json")
	if err != nil {
		return err
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	f.Write(b)

	return nil
}

func readToken() (string, error) {
	fp := ".cache/token.json"

	b, err := os.ReadFile(fp)
	if err != nil {
		return "", err
	}

	var authRespone AuthResponse
	err = json.Unmarshal(b, &authRespone)
	if err != nil {
		return "", err
	}

	return authRespone.Data.SessionToken, nil
}

func readAccountNumber() (string, error) {
	fp := ".cache/account.json"
	b, err := os.ReadFile(fp)
	if err != nil {
		return "", nil
	}

	var account AccountCache
	err = json.Unmarshal(b, &account)
	if err != nil {
		return "", nil
	}

	return account.Number, nil
}
