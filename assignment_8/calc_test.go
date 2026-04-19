package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, 5, Add(2, 3))
}

func TestDivideSuccess(t *testing.T) {
	result, err := Divide(10, 2)

	assert.NoError(t, err)
	assert.Equal(t, 5, result)
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)

	assert.Error(t, err)
}

func TestSubtractTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive positive", 5, 3, 2},
		{"positive zero", 5, 0, 5},
		{"negative positive", -5, 3, -8},
		{"negative negative", -5, -3, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
