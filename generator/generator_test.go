package generator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	g := Generator{}
	g.RegisterType(struct{ test int }{})
	_, err := g.Generate()
	if err != nil {
		t.Fatalf("Failed : %v", err)
	}
}
