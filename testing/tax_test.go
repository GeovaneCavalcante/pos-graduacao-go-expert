package tax

import "testing"

func TestCalculateTax(t *testing.T) {
	amount := 500.0
	expectedTax := 5.0

	result := CalculateTax(amount)

	if result != expectedTax {
		t.Errorf("Tax calculation incorrect. Expected: %f, Got: %f", expectedTax, result)
	}
}

func TestCaclculateTaxBtach(t *testing.T) {
	type calcTax struct {
		amount, expectedTax float64
	}

	table := []calcTax{
		{500, 5},
		{1000, 10},
		{2000, 10},
		{3000, 10},
		{0, 0},
	}

	for _, test := range table {
		result := CalculateTax(test.amount)

		if result != test.expectedTax {
			t.Errorf("Tax calculation incorrect. Expected: %f, Got: %f", test.expectedTax, result)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{500, 1000, 2000, 3000, 0}
	for _, amount := range seed {
		f.Add(amount)
	}
	f.Fuzz(func(t *testing.T, amount float64) {
		result := CalculateTax(amount)
		if amount <= 0 && result != 0 {
			t.Errorf("Tax calculation incorrect. Expected: %v, Got: %f", 0, result)
		}
		if amount > 20000 && result != 20 {
			t.Errorf("Tax calculation incorrect. Expected: %v, Got: %f", 0, result)
		}
	})
}
