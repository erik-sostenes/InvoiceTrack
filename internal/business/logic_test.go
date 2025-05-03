package business

import (
	"os"
	"path/filepath"
	"testing"
)

func BenchmarkTotalAmountCalculator(b *testing.B) {
	totalAmountCalculator, _ := NewTotalAmountCalculator()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		b.Fatal(err)
	}

	defaultXMLDir := filepath.Join(homeDir, "Documents", "xml_repo")

	for i := 0; i < b.N; i++ {
		_, _ = totalAmountCalculator.Calculate(defaultXMLDir, "xml")
	}
}
