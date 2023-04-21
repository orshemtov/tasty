package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TODO
// const maxSymbolSummaryBatchSize = 500

type MarketDataResponse struct {
	Data struct {
		Items []struct {
			Symbol                                 string    `json:"symbol"`
			ImpliedVolatilityIndex                 string    `json:"implied-volatility-index"`
			ImpliedVolatilityIndex5DayChange       string    `json:"implied-volatility-index-5-day-change"`
			ImpliedVolatilityIndexRank             string    `json:"implied-volatility-index-rank"`
			TosImpliedVolatilityIndexRank          string    `json:"tos-implied-volatility-index-rank"`
			TwImpliedVolatilityIndexRank           string    `json:"tw-implied-volatility-index-rank"`
			TosImpliedVolatilityIndexRankUpdatedAt time.Time `json:"tos-implied-volatility-index-rank-updated-at"`
			ImpliedVolatilityIndexRankSource       string    `json:"implied-volatility-index-rank-source"`
			ImpliedVolatilityPercentile            string    `json:"implied-volatility-percentile"`
			ImpliedVolatilityUpdatedAt             time.Time `json:"implied-volatility-updated-at"`
			LiquidityValue                         string    `json:"liquidity-value"`
			LiquidityRank                          string    `json:"liquidity-rank"`
			LiquidityRating                        int       `json:"liquidity-rating"`
			UpdatedAt                              time.Time `json:"updated-at"`
			OptionExpirationImpliedVolatilities    []struct {
				ExpirationDate    string `json:"expiration-date"`
				OptionChainType   string `json:"option-chain-type"`
				SettlementType    string `json:"settlement-type"`
				ImpliedVolatility string `json:"implied-volatility"`
			} `json:"option-expiration-implied-volatilities"`
			LiquidityRunningState struct {
				Sum       string    `json:"sum"`
				Count     int       `json:"count"`
				StartedAt time.Time `json:"started-at"`
				UpdatedAt time.Time `json:"updated-at"`
			} `json:"liquidity-running-state"`
			Beta                   string    `json:"beta"`
			BetaUpdatedAt          time.Time `json:"beta-updated-at"`
			CorrSpy3Month          string    `json:"corr-spy-3month"`
			DividendRatePerShare   string    `json:"dividend-rate-per-share"`
			AnnualDividendPerShare string    `json:"annual-dividend-per-share,omitempty"`
			DividendYield          string    `json:"dividend-yield"`
			DividendExDate         string    `json:"dividend-ex-date,omitempty"`
			DividendNextDate       string    `json:"dividend-next-date,omitempty"`
			DividendPayDate        string    `json:"dividend-pay-date,omitempty"`
			DividendUpdatedAt      time.Time `json:"dividend-updated-at,omitempty"`
			Earnings               struct {
				Visible            bool      `json:"visible"`
				ExpectedReportDate string    `json:"expected-report-date"`
				Estimated          bool      `json:"estimated"`
				TimeOfDay          string    `json:"time-of-day"`
				LateFlag           int       `json:"late-flag"`
				QuarterEndDate     string    `json:"quarter-end-date"`
				ActualEps          string    `json:"actual-eps"`
				ConsensusEstimate  string    `json:"consensus-estimate"`
				UpdatedAt          time.Time `json:"updated-at"`
			} `json:"earnings"`
			ListedMarket              string `json:"listed-market"`
			Lendability               string `json:"lendability"`
			BorrowRate                string `json:"borrow-rate"`
			MarketCap                 int64  `json:"market-cap"`
			ImpliedVolatility30Day    string `json:"implied-volatility-30-day"`
			HistoricalVolatility30Day string `json:"historical-volatility-30-day"`
			HistoricalVolatility60Day string `json:"historical-volatility-60-day"`
			HistoricalVolatility90Day string `json:"historical-volatility-90-day"`
			IvHv30DayDifference       string `json:"iv-hv-30-day-difference"`
		} `json:"items"`
	} `json:"data"`
}

func MarketMetrics(symbols []string) (*MarketDataResponse, error) {
	endpoint := fmt.Sprintf("/market-metrics?symbols=%s", strings.Join(symbols, ","))

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

	var r MarketDataResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}