package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccounts(t *testing.T) {
	r, err := Accounts()
	if err != nil {
		t.Errorf("Accounts() error = %v", err)
		return
	}

	got := r.Data.Items[0].Account.AccountNumber
	want := "[0-9A-Z]{8}"

	if !assert.Regexp(t, regexp.MustCompile(want), got) {
		t.Errorf("expected account number to match pattern: %s, got: %s", want, got)
		return
	}
}
