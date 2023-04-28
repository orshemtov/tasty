package tasty

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBalances(t *testing.T) {
	r, err := Balances()
	if err != nil {
		t.Errorf("Balances() error = %v", err)
		return
	}

	got, err := strconv.ParseFloat(r.Data.NetLiquidatingValue, 64)
	if err != nil {
		t.Errorf("could not parse net liquidation value, got = %s", r.Data.NetLiquidatingValue)
		return
	}
	want := 0.0

	if !assert.GreaterOrEqual(t, got, want) {
		t.Errorf("expected net liquidation value to be %f or bigger, got = %f", want, got)
		return
	}
}
