package parser

import (
	"errors"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"reflect"
	"testing"
	"time"
)

type ab_cdef struct {
	myint  int
	mytime time.Time
}

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
				name     int
				desc     int
				creation time.Time
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeInt},
					{Name: "desc", Type: ast.TypeInt},
					{Name: "creation", Type: ast.TypeDateTime},
				},
			},
		},
		{
			input: ab_cdef{},
			expected: ast.Table{
				Name: "ab_cdef",
				Fields: []ast.Field{
					{Name: "myint", Type: ast.TypeInt},
					{Name: "mytime", Type: ast.TypeDateTime},
				},
			},
		},
		{
			input: struct{ name string }{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeText},
				},
			},
		},
		{
			input: struct {
				name string `type:"varchar" size:"36"`
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeVarchar, Size: 36},
				},
			},
		},
		{
			input: struct {
				name string `type:"char" size:"36"`
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "name", Type: ast.TypeChar, Size: 36},
				},
			},
		},
		{
			input: struct {
				id int `primary:"true"`
			}{},
			expected: ast.Table{
				Name:         "anonym_1",
				PrimaryField: "id",
				Fields: []ast.Field{
					{Name: "id", Type: ast.TypeInt},
				},
			},
		},
		{
			input: struct {
				id int `primary:"true" autoincrement:"true"`
			}{},
			expected: ast.Table{
				Name:         "anonym_1",
				PrimaryField: "id",
				Fields: []ast.Field{
					{Name: "id", Type: ast.TypeInt, AutoIncrement: true},
				},
			},
		},
		{
			input: struct {
				id int `autoincrement:"true"`
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "id", Type: ast.TypeInt, AutoIncrement: true},
				},
			},
		},
		{
			input: struct {
				id int `type:"integer"`
			}{},
			expected: ast.Table{
				Name: "anonym_1",
				Fields: []ast.Field{
					{Name: "id", Type: ast.TypeInteger},
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
		{
			input: struct {
				name string `type:"varchar"`
			}{},
			expected: ErrTypeRequiredASize,
		},
		{
			input: struct {
				name string `type:"char"`
			}{},
			expected: ErrTypeRequiredASize,
		},
		{
			input: struct {
				id  int `primary:"true" autoincrement:"true"`
				id2 int `primary:"true" autoincrement:"true"`
			}{},
			expected: ErrSeveralPrimaryField,
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
