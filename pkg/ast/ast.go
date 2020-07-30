package ast

type Type struct {
	value string
}

func (t Type) String() string { return t.value }

var (
	TypeDecimal   = Type{"decimal"}
	TypeFloat     = Type{"float"}
	TypeDouble    = Type{"double"}
	TypeTinyint   = Type{"tinyint"}
	TypeSmallInt  = Type{"smallint"}
	TypeMediumInt = Type{"mediumint"}
	TypeInt       = Type{"int"}
	TypeBigInt    = Type{"bigint"}
	TypeChar      = Type{"char"}
	TypeVarchar   = Type{"varchar"}
	TypeDate      = Type{"date"}
	TypeDateTime  = Type{"datetime"}
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
