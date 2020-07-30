package generator

import (
	"database/sql"
	"github.com/canadadry/gostruct-to-sql/pkg/generator"
	"github.com/canadadry/gostruct-to-sql/pkg/parser"
	"strings"
)

type Generator struct {
	types    []interface{}
	Protocol string
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
		var query string
		switch g.Protocol {
		case "mysql":
			query, err = generator.MySql(ast)
		case "sqlite3":
			query, err = generator.Sqlite(ast)
		default:
			query, err = generator.Sqlite(ast)
		}
		if err != nil {
			return "", err
		}
		queries = append(queries, query)
	}
	return strings.Join(queries, "\n"), nil
}
