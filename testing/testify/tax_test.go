package testify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)

	assert.NoError(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(-1000.0)

	assert.Error(t, err)
	assert.Equal(t, 0.0, tax)
}
