package generator

import (
	"database/sql"
	"github.com/canadadry/gostruct-to-sql/pkg/generator"
	"github.com/canadadry/gostruct-to-sql/pkg/parser"
	"strings"
)

type Generator struct {
	types []interface{}
}

func (g *Generator) RegisterType(t interface{}) {
	g.types = append(g.types, t)
}

func (g *Generator) IsUpToDate(*sql.DB) bool {
	return false
}

func (g *Generator) Generate() (string, error) {
	queries := []string{}

	for _, t := range g.types {
		ast, err := parser.Parse(t)
		if err != nil {
			return "", err
		}
		query, err := generator.MySql(ast)
		if err != nil {
			return "", err
		}
		queries = append(queries, query)
	}
	return strings.Join(queries, "\n"), nil
}
