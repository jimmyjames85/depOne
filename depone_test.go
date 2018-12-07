package depone

import (
	"testing"
)

func TestOne(t *testing.T) {
	if "One" != OneString() {
		t.Fatal("should be One")
	}

}

func TestOneInt(t *testing.T) {
	if 1 != OneInt() {
		t.Fatal("should be 1")
	}

}
