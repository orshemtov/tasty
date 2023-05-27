package tasty

import (
	"encoding/json"
	"os"
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

func TestMarketMetricsManySymbols(t *testing.T) {
	data, err := os.ReadFile("testdata/symbols.json")
	if err != nil {
		t.Fatalf("could not read file, err = %v", err)
	}

	var symbols []string
	err = json.Unmarshal(data, &symbols)
	if err != nil {
		t.Fatalf("could not unmarshal symbols, err = %v", err)
	}

	r, err := MarketMetrics(symbols)
	if err != nil {
		t.Errorf("MarketMetrics() error = %v", err)
		return
	}

	want := 960 // some symbols are not returned

	assert.GreaterOrEqualf(t, len(r.Data.Items), want, "expected length to be greater than want = %v, got = %v", want, len(r.Data.Items))
}
