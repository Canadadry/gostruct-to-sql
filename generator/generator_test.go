package generator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	g := Generator{}
	err := g.RegisterType(struct{ test int }{})
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
	_, err = g.Generate()
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
}
