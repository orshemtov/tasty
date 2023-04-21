package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarketMetrics(t *testing.T) {
	symbols := []string{"AAPL", "GOOG"}
	r, err := MarketMetrics(symbols)
	if err != nil {
		t.Errorf("MarketMetrics() error = %v", err)
		return
	}

	var got []string
	for _, item := range r.Data.Items {
		got = append(got, item.Symbol)
	}
	want := symbols

	if !assert.Equal(t, want, got) {
		t.Errorf("want = %v, got = %v", want, got)
		return
	}
}
