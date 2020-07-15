package parser

import (
	"app/pkg/ast"
	"fmt"
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
		sqlType, err := convertGoTypeToMysql(f.Type.Kind())
		if err != nil {
			return t, fmt.Errorf("While parsing field %s : %w", f.Name, err)
		}
		fmt.Println(f.Type.Kind())
		t.Fields = append(t.Fields, ast.Field{
			Name: f.Name,
			Type: sqlType,
		})
	}

	return t, nil
}

func convertGoTypeToMysql(t reflect.Kind) (ast.Type, error) {
	switch t {
	case reflect.Int:
		return ast.TypeInt, nil
	}
	return ast.TypeInt, fmt.Errorf("%w : got %v", ErrUnknownType, t)
}
