package generator

import (
	"app/pkg/ast"
	"errors"
	"testing"
)

func TestGenerator(t *testing.T) {

	tests := []struct {
		input    ast.Table
		expected string
	}{
		{
			input: ast.Table{Name: "test"},
			expected: `CREATE TABLE test (
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
		{
			input: ast.Table{Name: "test2"},
			expected: `CREATE TABLE test2 (
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
		{
			input: ast.Table{
				Name: "test3",
				Fields: []ast.Field{
					{
						Name: "id",
						Type: ast.TypeInt,
					},
				},
			},
			expected: `CREATE TABLE test3 (
	id int,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
		{
			input: ast.Table{
				Name: "test4",
				Fields: []ast.Field{
					{
						Name: "id1",
						Type: ast.TypeInt,
					},
					{
						Name: "id2",
						Type: ast.TypeInt,
					},
				},
			},
			expected: `CREATE TABLE test4 (
	id1 int,
	id2 int,
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
	}

	for i, tt := range tests {
		result, err := MySql(tt.input)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}

		if result != tt.expected {
			t.Fatalf("[%d] generator return %s\n expected %s", i, result, tt.expected)
		}
	}
}

func TestTranslateType(t *testing.T) {
	tests := []struct {
		input    ast.Field
		expected string
	}{
		{
			input:    ast.Field{Name: "id1", Type: ast.TypeInt},
			expected: "int",
		},
		{
			input:    ast.Field{Name: "id1", Type: ast.TypeVarchar, Size: 10},
			expected: "varchar(10)",
		},
	}

	for i, tt := range tests {
		result, err := translateType(tt.input)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}

		if result != tt.expected {
			t.Fatalf("[%d] return %s\n expected %s", i, result, tt.expected)
		}
	}
}

func TestTranslateTypeError(t *testing.T) {
	tests := []struct {
		input    ast.Field
		expected error
	}{
		{
			input:    ast.Field{Type: ast.TypeChar},
			expected: errSizeOfFieldCannotBeZero,
		},
		{
			input:    ast.Field{Type: ast.TypeVarchar},
			expected: errSizeOfFieldCannotBeZero,
		},
	}

	for i, tt := range tests {
		_, err := translateType(tt.input)
		if err == nil {
			t.Fatalf("[%d] %#v sould have failed", i, tt.input)
		}

		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] return %v\n expected %v", i, err, tt.expected)
		}
	}
}
