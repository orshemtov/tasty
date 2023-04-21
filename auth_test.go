package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	r, err := Auth()
	if err != nil {
		t.Errorf("Auth() error = %v", err)
		return
	}

	want := username
	got := r.Data.User.Username

	if !assert.Equal(t, want, got) {
		t.Errorf("want = %s, got = %s", want, got)
		return
	}
}
