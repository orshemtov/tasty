package tasty

import (
	"encoding/json"
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

	respBody, err := doRequest(req)
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
