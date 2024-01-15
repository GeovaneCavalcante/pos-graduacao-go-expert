package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 5.0

	result := CalculateTax(amount)

	assert.Equal(t, expectedTax, result)
}
