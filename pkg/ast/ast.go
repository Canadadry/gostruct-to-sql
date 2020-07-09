package ast

const (
	TypeDecimal   = "decimal"
	TypeFloat     = "float"
	TypeDouble    = "double"
	TypeTinyint   = "tinyint"
	TypeSmallInt  = "smallint"
	TypeMediumInt = "mediumint"
	TypeInt       = "int"
	TypeBigInt    = "bigint"
	TypeChar      = "char"
	TypeVarchar   = "varchar"
	TypeDate      = "date"
	TypeDateTime  = "datetime"
)

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name          string
	Type          Type
	Size          uint
	Primary       bool
	AutoIncrement bool
	Nullable      bool
}

type Type string
