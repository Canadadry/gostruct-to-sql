package generator

import (
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"strings"
)

var (
	errSizeOfFieldCannotBeZero = fmt.Errorf("errSizeOfFieldCannotBeZero")
)

func MySql(t ast.Table) (string, error) {
	header := "CREATE TABLE " + t.Name + " (\n"
	footer := ") ENGINE=InnoDB DEFAULT CHARSET=utf8;"

	fields := make([]string, 0, len(t.Fields))
	for _, f := range t.Fields {
		typeName, err := translateType(f)
		if err != nil {
			return "", fmt.Errorf("Cannot translate field %#v : %w", f, err)
		}
		field := "\t" + f.Name + " " + typeName
		fields = append(fields, field)
	}
	content := strings.Join(fields, ",\n")
	if len(content) > 0 {
		content += "\n"
	}
	return header + content + footer, nil
}

func translateType(f ast.Field) (string, error) {
	if f.Type == ast.TypeChar || f.Type == ast.TypeVarchar {
		if f.Size == 0 {
			return "", errSizeOfFieldCannotBeZero
		}
		return fmt.Sprintf("%s(%d)", f.Type.String(), f.Size), nil
	}
	return f.Type.String(), nil
}
