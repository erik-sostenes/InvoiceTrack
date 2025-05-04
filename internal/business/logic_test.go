package business

import (
	"testing"
)

func BenchmarkTotalAmountCalculator(b *testing.B) {
	totalAmountCalculator, _ := NewTotalAmountCalculator()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = totalAmountCalculator.Calculate("../../test_files/", "xml")
	}
}
