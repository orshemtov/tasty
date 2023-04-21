package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type AccountsResponse struct {
	Data struct {
		Items []struct {
			AuthorityLevel string `json:"authority-level"`
			Account        struct {
				AccountNumber         string    `json:"account-number"`
				ExternalID            string    `json:"external-id"`
				OpenedAt              time.Time `json:"opened-at"`
				Nickname              string    `json:"nickname"`
				AccountTypeName       string    `json:"account-type-name"`
				DayTraderStatus       bool      `json:"day-trader-status"`
				IsClosed              bool      `json:"is-closed"`
				IsFirmError           bool      `json:"is-firm-error"`
				IsFirmProprietary     bool      `json:"is-firm-proprietary"`
				IsFuturesApproved     bool      `json:"is-futures-approved"`
				IsTestDrive           bool      `json:"is-test-drive"`
				MarginOrCash          string    `json:"margin-or-cash"`
				IsForeign             bool      `json:"is-foreign"`
				FundingDate           string    `json:"funding-date"`
				InvestmentObjective   string    `json:"investment-objective"`
				LiquidityNeeds        string    `json:"liquidity-needs"`
				RiskTolerance         string    `json:"risk-tolerance"`
				InvestmentTimeHorizon string    `json:"investment-time-horizon"`
				FuturesAccountPurpose string    `json:"futures-account-purpose"`
				SuitableOptionsLevel  string    `json:"suitable-options-level"`
				CreatedAt             time.Time `json:"created-at"`
			} `json:"account,omitempty"`
		} `json:"items"`
	} `json:"data"`
	Context string `json:"context"`
}

type AccountCache struct {
	Number string `json:"number"`
}

func Accounts() (*AccountsResponse, error) {
	endpoint := "/customers/me/accounts"

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

	var r AccountsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
