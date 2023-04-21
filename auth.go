package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type AuthPayload struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Data struct {
		User struct {
			Email      string `json:"email"`
			Username   string `json:"username"`
			ExternalID string `json:"external-id"`
		} `json:"user"`
		SessionToken string `json:"session-token"`
	} `json:"data"`
	Context string `json:"context"`
}

func Auth() (*AuthResponse, error) {
	endpoint := "/sessions"

	u, err := url.Parse(baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	payload := AuthPayload{
		Username: username,
		Password: password,
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(payloadJSON))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		panic(resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r AuthResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
