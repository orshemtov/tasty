package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicWatchlists(t *testing.T) {
	r, err := PublicWatchlists()
	if err != nil {
		t.Errorf("PublicWatchlists() error = %v", err)
		return
	}

	var got []string
	for _, wl := range r.Data.Items {
		got = append(got, wl.GroupName)
	}

	for _, group := range []string{"Market Indices", "Liquidity", "Earnings", "Futures"} {
		if !assert.Contains(t, got, group) {
			t.Errorf("expected to find group = %s in %v", group, got)
			return
		}
	}
}
