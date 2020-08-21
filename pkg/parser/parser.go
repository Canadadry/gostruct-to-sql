package parser

import (
	"fmt"
	"github.com/canadadry/gostruct-to-sql/pkg/ast"
	"reflect"
	"strconv"
)

var (
	ErrNotAStruct          = fmt.Errorf("Can only parse struct type")
	ErrUnknownType         = fmt.Errorf("Cannot convert unknown go type")
	ErrTypeRequiredASize   = fmt.Errorf("This type require a size annotation")
	ErrSeveralPrimaryField = fmt.Errorf("Several primary field ")
)

const (
	tagType          = "type"
	tagSize          = "size"
	tagPrimary       = "primary"
	tagAutoIncrement = "autoincrement"
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
		tags := tOfv.Field(i).Tag
		sqlType, sqlSize, err := convertGoTypeToSqlType(f.Type, tags)
		if err != nil {
			return t, fmt.Errorf("While parsing field %s : %w", f.Name, err)
		}

		t.Fields = append(t.Fields, ast.Field{
			Name:          f.Name,
			Type:          sqlType,
			Size:          sqlSize,
			AutoIncrement: isAutoIncrement(tags),
		})
		if isPrimaryKey(tags) {
			if len(t.PrimaryField) > 0 {
				return t, ErrSeveralPrimaryField
			}
			t.PrimaryField = f.Name
		}
	}

	return t, nil
}

func convertGoTypeToSqlType(t reflect.Type, tags reflect.StructTag) (ast.Type, uint, error) {
	switch t.Kind() {
	case reflect.Int:
		tt, _ := tags.Lookup(tagType)
		switch tt {
		case "integer":
			return ast.TypeInteger, 0, nil
		}
		return ast.TypeInt, 0, nil
	case reflect.String:
		tt, _ := tags.Lookup(tagType)
		switch tt {
		case "varchar":
			size, err := readSizeOf(tags)
			if err != nil {
				return ast.TypeVarchar, 0, fmt.Errorf("Cannot find size of %v : %w", t, err)
			}
			return ast.TypeVarchar, size, nil
		case "char":
			size, err := readSizeOf(tags)
			if err != nil {
				return ast.TypeChar, 0, fmt.Errorf("Cannot find size of %v : %w", t, err)
			}
			return ast.TypeChar, size, nil
		}

		return ast.TypeText, 0, nil
	case reflect.Struct:
		if t.PkgPath() == "time" && t.Name() == "Time" {
			return ast.TypeDateTime, 0, nil
		}
		return ast.TypeInt, 0, nil
	}
	return ast.Type{}, 0, fmt.Errorf("%w : got %v", ErrUnknownType, t)
}

func readSizeOf(tags reflect.StructTag) (uint, error) {

	ts, ok := tags.Lookup(tagSize)
	if !ok {
		return 0, ErrTypeRequiredASize
	}
	if len(ts) == 0 {
		return 0, ErrTypeRequiredASize
	}

	size, err := strconv.ParseUint(ts, 10, 32)
	if err != nil {
		return 0, ErrTypeRequiredASize
	}
	return uint(size), nil
}

func isPrimaryKey(tags reflect.StructTag) bool {
	_, ok := tags.Lookup(tagPrimary)
	return ok
}

func isAutoIncrement(tags reflect.StructTag) bool {
	_, ok := tags.Lookup(tagAutoIncrement)
	return ok
}
