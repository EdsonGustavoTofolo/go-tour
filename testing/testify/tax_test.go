package testify

import (
	"errors"
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

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &TaxRepositoryMock{}
	repository.On("SaveTax", 10.0).Return(nil).Once()
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repository)
	assert.NoError(t, err)

	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err)

	repository.AssertExpectations(t)
	repository.AssertCalled(t, "SaveTax", 10.0)
	repository.AssertNumberOfCalls(t, "SaveTax", 2)
}
