package parser

import (
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"reflect"
)

var (
	ErrNotAStruct  = fmt.Errorf("Can only parse struct type")
	ErrUnknownType = fmt.Errorf("Cannot convert unknown go type")
)

func Parse(v interface{}) (ast.Table, error) {
	t := ast.Table{}
	tOfv := reflect.TypeOf(v)

	if tOfv.Kind() != reflect.Struct {
		return t, fmt.Errorf("%w : got %s", ErrNotAStruct, tOfv.Kind())
	}

	t.Name = tOfv.Name()
	if len(t.Name) == 0 {
		t.Name = "anonym_1"
	}

	for i := 0; i < tOfv.NumField(); i++ {
		f := tOfv.Field(i)
		sqlType, err := convertGoTypeToSqlType(f.Type)
		if err != nil {
			return t, fmt.Errorf("While parsing field %s : %w", f.Name, err)
		}
		// fmt.Println(f.Type.Kind())
		t.Fields = append(t.Fields, ast.Field{
			Name: f.Name,
			Type: sqlType,
		})
	}

	return t, nil
}

func convertGoTypeToSqlType(t reflect.Type) (ast.Type, error) {
	switch t.Kind() {
	case reflect.Int:
		return ast.TypeInt, nil
	case reflect.Struct:
		if t.PkgPath() == "time" && t.Name() == "Time" {
			return ast.TypeDateTime, nil
		}
		return ast.TypeInt, nil
	}
	return ast.Type{}, fmt.Errorf("%w : got %v", ErrUnknownType, t)
}
