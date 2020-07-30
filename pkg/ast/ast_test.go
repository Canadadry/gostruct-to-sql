package ast

import (
	"testing"
)

func TestString(t *testing.T) {
	input := "test"
	ty := Type{input}

	if ty.String() != input {
		t.Fatalf("failed exp %s, got %s", input, ty.String())
	}
}
