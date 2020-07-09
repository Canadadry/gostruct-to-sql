package ast

const (
	TypeInt     = "int"
	TypeChar    = "char"
	TypeDecimal = "decimal"
	TypeVarchar = "varchar"
	TypeTinyint = "tinyint"
)

type Table struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name          string
	Type          Type
	Size1         uint
	Size2         uint
	Primary       bool
	AutoIncrement bool
	Nullable      bool
}

type Type string
