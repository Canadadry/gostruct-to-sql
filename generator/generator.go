package generator

import (
	"database/sql"
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"github.com/canadadry/gostruct-to-sql/pkg/generator"
	"github.com/canadadry/gostruct-to-sql/pkg/parser"
	"strings"
)

type Generator struct {
	types    map[string]ast.Table
	Protocol string
}

func (g *Generator) RegisterType(t interface{}) error {
	p, err := parser.Parse(t)
	if err != nil {
		return err
	}
	if g.types == nil {
		g.types = map[string]ast.Table{}
	}
	g.types[p.Name] = p
	return nil
}

const tableExistQuery = `select * from %s;`

func doesTableExist(db *sql.DB, name string) error {
	rows, err := db.Query(fmt.Sprintf(tableExistQuery, name))
	if err != nil {
		return fmt.Errorf("query failed :%w", err)
	}
	if err := rows.Close(); err != nil {
		return fmt.Errorf("row close failed :%w", err)
	}
	return nil
}

func (g *Generator) IsUpToDate(db *sql.DB) bool {
	for k, _ := range g.types {
		err := doesTableExist(db, k)
		if err != nil {
			return false
		}
	}
	return true
}

func (g *Generator) Generate() (string, error) {
	queries := []string{}

	for _, t := range g.types {
		var query string
		var err error
		switch g.Protocol {
		case "mysql":
			query, err = generator.MySql(t)
		case "sqlite3":
			query, err = generator.Sqlite(t)
		default:
			query, err = generator.Sqlite(t)
		}
		if err != nil {
			return "", err
		}
		queries = append(queries, query)
	}
	return strings.Join(queries, "\n"), nil

}
