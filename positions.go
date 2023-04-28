package tasty

import (
	"encoding/json"
	"fmt"
	"io"
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

	var r PositionsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
