package native

import "testing"

/*
	HOW TO RUN TESTS?
	Inside the directory that have the tests files run the following command:
	> go test -v
	> go test

	HOW TO RUN COVERAGE TESTS?
	> go test -coverprofile=coverage.out
	> go tool cover -html=coverage.out

	HOW TO RUN BENCHMARK?
	Initialize func with name Benchmark
	> go test -bench=.
	> go test -bench=. -count=10
	> go test -bench=. -benchmem
	> go help test

	HOW TO RUN FUZZING/MUTATION TESTS?
	> go test -fuzz=.
*/

func TestCalculateTax(t *testing.T) {
	expectedTax := 5.0
	amount := 500.0

	tax := CalculateTax(amount)

	if tax != expectedTax {
		t.Errorf("Expected %f but got %f", expectedTax, tax)
	}
}

func TestCalculateTaxBatch(t *testing.T) {
	type calcTax struct {
		amount, expect float64
	}

	table := []calcTax{
		{500.0, 5.0},
		{999.9, 5.0},
		{1000.0, 10.0},
		{1500.0, 10.0},
	}

	for _, item := range table {
		tax := CalculateTax(item.amount)
		if tax != item.expect {
			t.Errorf("Expected %f but got %f", item.expect, tax)
		}
	}
}

func BenchmarkCalculateTax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax(500.0)
	}
}

func BenchmarkCalculateTax2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateTax2(500.0)
	}
}

func FuzzCalculateTax(f *testing.F) {
	seed := []float64{-1.0, 0.0, 500.0, 1000.0, 1500.0}

	for _, amount := range seed {
		f.Add(amount)
	}

	f.Fuzz(func(t *testing.T, amount float64) {
		tax := CalculateTax(amount)
		if amount <= 0 && tax != 0 {
			t.Errorf("Received %f but expect 0", tax)
		}
	})
}
