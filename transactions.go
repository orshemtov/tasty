package tasty

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type TransactionsResponse struct {
	Data struct {
		Items []struct {
			ID                               int       `json:"id"`
			AccountNumber                    string    `json:"account-number"`
			TransactionType                  string    `json:"transaction-type"`
			TransactionSubType               string    `json:"transaction-sub-type"`
			Description                      string    `json:"description"`
			ExecutedAt                       time.Time `json:"executed-at"`
			TransactionDate                  string    `json:"transaction-date"`
			Value                            string    `json:"value"`
			ValueEffect                      string    `json:"value-effect"`
			NetValue                         string    `json:"net-value"`
			NetValueEffect                   string    `json:"net-value-effect"`
			IsEstimatedFee                   bool      `json:"is-estimated-fee"`
			Symbol                           string    `json:"symbol,omitempty"`
			InstrumentType                   string    `json:"instrument-type,omitempty"`
			UnderlyingSymbol                 string    `json:"underlying-symbol,omitempty"`
			Action                           string    `json:"action,omitempty"`
			Quantity                         string    `json:"quantity,omitempty"`
			Price                            string    `json:"price,omitempty"`
			RegulatoryFees                   string    `json:"regulatory-fees,omitempty"`
			RegulatoryFeesEffect             string    `json:"regulatory-fees-effect,omitempty"`
			ClearingFees                     string    `json:"clearing-fees,omitempty"`
			ClearingFeesEffect               string    `json:"clearing-fees-effect,omitempty"`
			Commission                       string    `json:"commission,omitempty"`
			CommissionEffect                 string    `json:"commission-effect,omitempty"`
			ProprietaryIndexOptionFees       string    `json:"proprietary-index-option-fees,omitempty"`
			ProprietaryIndexOptionFeesEffect string    `json:"proprietary-index-option-fees-effect,omitempty"`
			ExtExchangeOrderNumber           string    `json:"ext-exchange-order-number,omitempty"`
			ExtGlobalOrderNumber             int       `json:"ext-global-order-number,omitempty"`
			ExtGroupID                       string    `json:"ext-group-id,omitempty"`
			ExtGroupFillID                   string    `json:"ext-group-fill-id,omitempty"`
			ExtExecID                        string    `json:"ext-exec-id,omitempty"`
			ExecID                           string    `json:"exec-id,omitempty"`
			Exchange                         string    `json:"exchange,omitempty"`
			OrderID                          int       `json:"order-id,omitempty"`
		} `json:"items"`
	} `json:"data"`
	Context    string `json:"context"`
	Pagination struct {
		PerPage            int         `json:"per-page"`
		PageOffset         int         `json:"page-offset"`
		ItemOffset         int         `json:"item-offset"`
		TotalItems         int         `json:"total-items"`
		TotalPages         int         `json:"total-pages"`
		CurrentItemCount   int         `json:"current-item-count"`
		PreviousLink       interface{} `json:"previous-link"`
		NextLink           interface{} `json:"next-link"`
		PagingLinkTemplate interface{} `json:"paging-link-template"`
	} `json:"pagination"`
}

func Transactions(startDate *time.Time, endDate *time.Time, pageOffset *int) (*TransactionsResponse, error) {
	endpoint := fmt.Sprintf("/accounts/%s/transactions", accountNumber)

	u, err := url.Parse(baseUrl + endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()

	if startDate != nil {
		q.Set("start-date", startDate.Format("2006-01-02"))
	}

	if endDate != nil {
		q.Set("end-date", endDate.Format("2006-01-02"))
	}

	if pageOffset != nil {
		q.Set("page-offset", strconv.Itoa(*pageOffset))
	}

	u.RawQuery = q.Encode()

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

	var r TransactionsResponse
	err = json.Unmarshal(respBody, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
