package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRateSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"rate":500}`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	rate, err := service.GetRate("USD", "KZT")

	assert.NoError(t, err)
	assert.Equal(t, 500.0, rate)
}

func TestGetRateError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"invalid pair"}`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("AAA", "BBB")

	assert.Error(t, err)
}
