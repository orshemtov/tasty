package tasty

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

type WatchlistsResponse struct {
	Data struct {
		Items []struct {
			Name             string `json:"name"`
			WatchlistEntries []struct {
				Symbol         string `json:"symbol"`
				InstrumentType string `json:"instrument-type,omitempty"`
			} `json:"watchlist-entries"`
			OrderIndex int `json:"order-index"`
		} `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

func Watchlists() (*WatchlistsResponse, error) {
	endpoint := "/watchlists"

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, endpoint)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	respBody, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var r WatchlistsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
