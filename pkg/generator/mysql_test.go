package generator

import (
	"app/pkg/ast"
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
