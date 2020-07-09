package generator

import (
	"app/pkg/ast"
	"strings"
)

func MySql(t ast.Table) (string, error) {
	header := "CREATE TABLE " + t.Name + " (\n"
	footer := ") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	fields := []string{}
	for _, f := range t.Fields {
		field := "\t" + f.Name + " " + string(f.Type) + ",\n"
		fields = append(fields, field)
	}
	return header + strings.Join(fields, "") + footer, nil
}
