package tasty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PositionsResponse struct {
	// TODO
}

func Positions() (*PositionsResponse, error) {
	endpoint := fmt.Sprintf("/accounts/%s/positions", accountNumber)

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

	var r PositionsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
