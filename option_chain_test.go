package tasty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionChain(t *testing.T) {
	symbol := "AAPL"
	r, err := OptionChain(symbol)
	if err != nil {
		t.Errorf("OptionChain() error = %v", err)
		return
	}

	got := len(r.Data.Items[0].Expirations)
	want := 1

	if !assert.GreaterOrEqual(t, got, want) {
		t.Errorf("got = %v, want = %v", got, want)
		return
	}
}
