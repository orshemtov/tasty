package tasty

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type BalancesResponse struct {
	Data struct {
		AccountNumber                      string    `json:"account-number"`
		CashBalance                        string    `json:"cash-balance"`
		LongEquityValue                    string    `json:"long-equity-value"`
		ShortEquityValue                   string    `json:"short-equity-value"`
		LongDerivativeValue                string    `json:"long-derivative-value"`
		ShortDerivativeValue               string    `json:"short-derivative-value"`
		LongFuturesValue                   string    `json:"long-futures-value"`
		ShortFuturesValue                  string    `json:"short-futures-value"`
		LongFuturesDerivativeValue         string    `json:"long-futures-derivative-value"`
		ShortFuturesDerivativeValue        string    `json:"short-futures-derivative-value"`
		LongMargineableValue               string    `json:"long-margineable-value"`
		ShortMargineableValue              string    `json:"short-margineable-value"`
		MarginEquity                       string    `json:"margin-equity"`
		EquityBuyingPower                  string    `json:"equity-buying-power"`
		DerivativeBuyingPower              string    `json:"derivative-buying-power"`
		DayTradingBuyingPower              string    `json:"day-trading-buying-power"`
		FuturesMarginRequirement           string    `json:"futures-margin-requirement"`
		AvailableTradingFunds              string    `json:"available-trading-funds"`
		MaintenanceRequirement             string    `json:"maintenance-requirement"`
		MaintenanceCallValue               string    `json:"maintenance-call-value"`
		RegTCallValue                      string    `json:"reg-t-call-value"`
		DayTradingCallValue                string    `json:"day-trading-call-value"`
		DayEquityCallValue                 string    `json:"day-equity-call-value"`
		NetLiquidatingValue                string    `json:"net-liquidating-value"`
		CashAvailableToWithdraw            string    `json:"cash-available-to-withdraw"`
		DayTradeExcess                     string    `json:"day-trade-excess"`
		PendingCash                        string    `json:"pending-cash"`
		PendingCashEffect                  string    `json:"pending-cash-effect"`
		LongCryptocurrencyValue            string    `json:"long-cryptocurrency-value"`
		ShortCryptocurrencyValue           string    `json:"short-cryptocurrency-value"`
		CryptocurrencyMarginRequirement    string    `json:"cryptocurrency-margin-requirement"`
		UnsettledCryptocurrencyFiatAmount  string    `json:"unsettled-cryptocurrency-fiat-amount"`
		UnsettledCryptocurrencyFiatEffect  string    `json:"unsettled-cryptocurrency-fiat-effect"`
		ClosedLoopAvailableBalance         string    `json:"closed-loop-available-balance"`
		EquityOfferingMarginRequirement    string    `json:"equity-offering-margin-requirement"`
		LongBondValue                      string    `json:"long-bond-value"`
		BondMarginRequirement              string    `json:"bond-margin-requirement"`
		SnapshotDate                       string    `json:"snapshot-date"`
		RegTMarginRequirement              string    `json:"reg-t-margin-requirement"`
		FuturesOvernightMarginRequirement  string    `json:"futures-overnight-margin-requirement"`
		FuturesIntradayMarginRequirement   string    `json:"futures-intraday-margin-requirement"`
		MaintenanceExcess                  string    `json:"maintenance-excess"`
		PendingMarginInterest              string    `json:"pending-margin-interest"`
		EffectiveCryptocurrencyBuyingPower string    `json:"effective-cryptocurrency-buying-power"`
		UpdatedAt                          time.Time `json:"updated-at"`
	} `json:"data"`
	Context string `json:"context"`
}

func Balances() (*BalancesResponse, error) {
	endpoint := fmt.Sprintf("/accounts/%s/balances", accountNumber)

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

	var r BalancesResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
