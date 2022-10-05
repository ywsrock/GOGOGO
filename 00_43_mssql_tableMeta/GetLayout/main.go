package main

import (
	"getlayout/export"
	"getlayout/ini"
	"getlayout/model"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	outMode = 2 // 1: text template 2:standart output
)

func main() {
	// github.com/denisenkom/go-mssqldb
	dsn := ini.GetDbString()
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalln(err)
	}

	var results []model.Result
	// Get the table meta info
	sql := `SELECT top(1)
        O.name AS TableName,
        C.name AS ColumnName,
		T.name AS DataType,
		c.MaxLength,
		c.Precision,
		c.Scale
		FROM sys.objects AS o
		INNER JOIN sys.columns AS c ON o.object_id = c.object_id
		INNER JOIN sys.types AS t ON c.user_type_id = t.user_type_id
		WHERE type = 'U'`

	// Excute SQL Fetch table meta info
	db.Raw(sql).Scan(&results)
	switch outMode {
	case 1:
		// Print to text template
		export.OutputTextTemp(&results)
	case 2:
		// Print to Standard output
		export.OutputStdout(results)
	}
}
