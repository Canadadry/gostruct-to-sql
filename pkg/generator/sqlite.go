package generator

import (
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"strings"
)

func Sqlite(t ast.Table) (string, error) {
	header := "CREATE TABLE " + t.Name + " (\n"
	footer := ");"

	fields := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		typeName, err := translateType(f)
		if err != nil {
			return "", fmt.Errorf("Cannot translate field %#v : %w", f, err)
		}
		field := "\t" + f.Name + " " + typeName + ",\n"
		fields = append(fields, field)
	}
	return header + strings.Join(fields, "") + footer, nil
}
