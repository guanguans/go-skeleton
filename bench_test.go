package skeleton

import "testing"

// go test -bench=. -benchmem ./...
func BenchmarkReturnTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReturnTrue()
	}
}
