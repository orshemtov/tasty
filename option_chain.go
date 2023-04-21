package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type OptionChainResponse struct {
	Data struct {
		Items []struct {
			UnderlyingSymbol  string `json:"underlying-symbol"`
			RootSymbol        string `json:"root-symbol"`
			OptionChainType   string `json:"option-chain-type"`
			SharesPerContract int    `json:"shares-per-contract"`
			TickSizes         []struct {
				Value     string `json:"value"`
				Threshold string `json:"threshold,omitempty"`
			} `json:"tick-sizes"`
			Deliverables []struct {
				ID              int    `json:"id"`
				RootSymbol      string `json:"root-symbol"`
				DeliverableType string `json:"deliverable-type"`
				Description     string `json:"description"`
				Amount          string `json:"amount"`
				Symbol          string `json:"symbol"`
				InstrumentType  string `json:"instrument-type"`
				Percent         string `json:"percent"`
			} `json:"deliverables"`
			Expirations []struct {
				ExpirationType   string `json:"expiration-type"`
				ExpirationDate   string `json:"expiration-date"`
				DaysToExpiration int    `json:"days-to-expiration"`
				SettlementType   string `json:"settlement-type"`
				Strikes          []struct {
					StrikePrice string `json:"strike-price"`
					Call        string `json:"call"`
					Put         string `json:"put"`
				} `json:"strikes"`
			} `json:"expirations"`
		} `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

func OptionChain(symbol string) (*OptionChainResponse, error) {
	endpoint := fmt.Sprintf("/option-chains/%s/nested", symbol)

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

	var r OptionChainResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
