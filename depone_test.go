package depone

import (
	"testing"
)

func TestOne(t *testing.T) {
	if "One" != One() {
		t.Fatal("should be One")
	}

}
