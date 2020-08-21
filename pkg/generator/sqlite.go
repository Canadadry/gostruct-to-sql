package generator

import (
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"strings"
)

var (
	ErrAutoIncrement = fmt.Errorf("Cannot have AutoIncrement without PrimaryKey and Integer")
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
			if f.Name == t.PrimaryField && typeName == ast.TypeInteger.String() {
				field += " PRIMARY KEY AUTOINCREMENT"
			} else {
				return "", ErrAutoIncrement
			}
		}
		fields = append(fields, field)
	}
	content := strings.Join(fields, ",\n")
	if len(content) > 0 {
		content += "\n"
	}
	return header + content + footer, nil
}
