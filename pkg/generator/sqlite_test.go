package generator

import (
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
				},
			},
			expected: `CREATE TABLE test4 (
	id1 int,
	id2 int
);`,
		},
	}

	for i, tt := range tests {
		result, err := Sqlite(tt.input)
		if err != nil {
			t.Fatalf("[%d] failed %v", i, err)
		}

		if result != tt.expected {
			t.Fatalf("[%d] generator return %s\n expected %s", i, result, tt.expected)
		}
	}
}
