package parser

import (
	"errors"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {

	tests := []struct {
		input    interface{}
		expected ast.Table
	}{
		{
			input:    struct{}{},
			expected: ast.Table{Name: "anonym_1"},
		},
		{
			input: struct {
				name int
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeInt},
				},
			},
		},
		{
			input: struct {
				name int
				desc int
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeInt},
					{Name: "desc", Type: ast.TypeInt},
				},
			},
		},
	}

	for i, tt := range tests {
		result, err := Parse(tt.input)
		if err != nil {
			t.Fatalf("failed : %v", err)
		}

		if !reflect.DeepEqual(result, tt.expected) {
			t.Fatalf("[%d] generator return %#v\n expected %#v", i, result, tt.expected)
		}
	}
}

func TestParserError(t *testing.T) {

	tests := []struct {
		input    interface{}
		expected error
	}{
		{
			input:    0,
			expected: ErrNotAStruct,
		},
		{
			input: struct {
				name bool
			}{},
			expected: ErrUnknownType,
		},
	}

	for i, tt := range tests {
		_, err := Parse(tt.input)
		if err == nil {
			t.Fatalf("[%d] should have failed", i)
		}

		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] return %v\n expected %v", i, err, tt.expected)
		}
	}
}
