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
		field := "\t" + f.Name + " " + typeName
		if f.AutoIncrement {
			field += " AUTOINCREMENT"
		}
		fields = append(fields, field)
	}
	if len(t.PrimaryField) > 0 {
		fields = append(fields, "\tPRIMARY KEY ("+t.PrimaryField+")")
	}
	content := strings.Join(fields, ",\n")
	if len(content) > 0 {
		content += "\n"
	}
	return header + content + footer, nil
}
