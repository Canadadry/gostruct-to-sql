package generator

import (
	"errors"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"testing"
)

func TestGeneratorSqlite(t *testing.T) {

	tests := []struct {
		input    ast.Table
		expected string
	}{
		{
			input: ast.Table{Name: "test"},
			expected: `CREATE TABLE test (
);`,
		},
		{
			input: ast.Table{Name: "test2"},
			expected: `CREATE TABLE test2 (
);`,
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
	id int
);`,
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
					{
						Name: "creation",
						Type: ast.TypeDateTime,
					},
				},
			},
			expected: `CREATE TABLE test4 (
	id1 int,
	id2 int,
	creation datetime
);`,
		},
		{
			input: ast.Table{
				Name: "test5",
				Fields: []ast.Field{
					{
						Name: "name",
						Type: ast.TypeText,
					},
				},
			},
			expected: `CREATE TABLE test5 (
	name text
);`,
		},
		{
			input: ast.Table{
				Name:         "test7",
				PrimaryField: "name",
				Fields: []ast.Field{
					{
						Name:          "name",
						Type:          ast.TypeInteger,
						AutoIncrement: true,
					},
				},
			},
			expected: `CREATE TABLE test7 (
	name integer PRIMARY KEY AUTOINCREMENT
);`,
		},
	}

	for i, tt := range tests {
		result, err := Sqlite(tt.input)
		if err != nil {
			t.Fatalf("[%d-%s] failed %v", i, tt.input.Name, err)
		}

		if result != tt.expected {
			t.Fatalf("[%d] generator return %s\n expected %s", i, result, tt.expected)
		}
	}
}

func TestGeneratorSqliteError(t *testing.T) {

	tests := []struct {
		input    ast.Table
		expected error
	}{
		{
			input: ast.Table{
				Name: "test6",
				Fields: []ast.Field{
					{
						Name:          "name",
						Type:          ast.TypeInteger,
						AutoIncrement: true,
					},
				},
			},
			expected: ErrAutoIncrement,
		},
		{
			input: ast.Table{
				Name:         "test6",
				PrimaryField: "name",
				Fields: []ast.Field{
					{
						Name:          "name",
						Type:          ast.TypeInt,
						AutoIncrement: true,
					},
				},
			},
			expected: ErrAutoIncrement,
		},
	}

	for i, tt := range tests {
		_, err := Sqlite(tt.input)
		if err == nil {
			t.Fatalf("[%d] shoul d have failed", i)
		}
		if !errors.Is(err, tt.expected) {
			t.Fatalf("[%d] return %v\n expected %v", i, err, tt.expected)
		}
	}
}
