// Description: Modern benchmarks with testing.B.Loop (Go 1.24)
// Tags: testing, benchmark, B.Loop, strings, Go 1.24
package benchmarkloop

import (
	"strings"
	"testing"
)

func BenchmarkJoin(b *testing.B) {
	parts := []string{"range", "over", "function", "iterator"}

	for b.Loop() {
		strings.Join(parts, "-")
	}
}
