package skeleton

import (
	"testing"
)

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestReturnTrue(t *testing.T) {
	if !ReturnTrue() {
		t.Errorf("Test failed")
	}
}
