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
	TypeInteger   = Type{"integer"}
	TypeBigInt    = Type{"bigint"}
	TypeChar      = Type{"char"}
	TypeVarchar   = Type{"varchar"}
	TypeText      = Type{"text"}
	TypeDate      = Type{"date"}
	TypeDateTime  = Type{"datetime"}
)

type Table struct {
	Name         string
	Fields       []Field
	PrimaryField string
}

type Field struct {
	Name          string
	Type          Type
	Size          uint
	AutoIncrement bool
	Nullable      bool
}
