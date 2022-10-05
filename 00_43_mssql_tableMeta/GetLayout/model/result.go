package model

type Result struct {
	TableName  string
	ColumnName string
	DataType   string
	MaxLength  string
	Precision  string
	Scale      string
}

type ForeignKey struct {
	TableName  string
	ColumnName string
}

type PrimaryKey struct {
	TableName  string
	ColumnName string
}
