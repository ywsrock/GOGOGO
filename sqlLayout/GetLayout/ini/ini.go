package ini

import (
	"fmt"
	"regexp"

	"github.com/zieckey/goini"
)

var ini *goini.INI

var (
	filename     string = "XXXXXX.ini"
	dbServerName        = ""
	dbFileName          = ""
	dbUserName          = ""
	dbPasswd            = ""
)

func GetDbString() (dsn string) {
	ini = goini.New()
	err := ini.ParseFile(filename)
	if err != nil {
		fmt.Printf("parse INI file %v failed : %v\n", filename, err.Error())
		return
	}
	v, _ := ini.SectionGet("DB", "DB_SERVER_NAME")
	v0, _ := ini.SectionGet("DB", "DB_FILE_NAME")
	v1, _ := ini.SectionGet("DB", "DB_USER_ID")
	v2, _ := ini.SectionGet("DB", "DB_PASSWORD")
	dbServerName = replaceStr(v)
	dbFileName = replaceStr(v0)
	dbUserName = replaceStr(v1)
	dbPasswd = replaceStr(v2)
	dsn = fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbUserName, dbPasswd, dbServerName, dbFileName)
	return
}

func replaceStr(src string) (dst string) {
	re := regexp.MustCompile(`\s+.*`)
	dst = re.ReplaceAllString(src, "")
	return dst
}
