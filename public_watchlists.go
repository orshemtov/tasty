package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type PublicWatchlistsResponse struct {
	Data struct {
		Items []struct {
			Name             string `json:"name"`
			WatchlistEntries []struct {
				Symbol         string `json:"symbol"`
				InstrumentType string `json:"instrument-type"`
			} `json:"watchlist-entries"`
			GroupName  string `json:"group-name"`
			OrderIndex int    `json:"order-index"`
		} `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

func PublicWatchlists() (*PublicWatchlistsResponse, error) {
	endpoint := "/public-watchlists"

	u, err := url.Parse(baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r PublicWatchlistsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
