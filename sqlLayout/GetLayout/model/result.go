package model

/**
 * Table meta Info struct
 *
**/
type Result struct {
	TableName  string
	ColumnName string
	DataType   string
	Max_Length string
	Precision  string
	Scale      string
}

/**
 *
 * Foreign Key
**/
type ForeignKey struct {
	TableName  string
	ColumnName string
}

/**
 *
 *　Primary Key
**/
type PrimaryKey struct {
	TableName  string
	ColumnName string
}
