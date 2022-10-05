package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pageredder/model"
)

var db *gorm.DB

func DBInit() {
	db1, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:passw0rd@tcp(127.0.0.1:3306)/wordlist?charset=utf8mb4&parseTime=True&loc=Local", // data source name
		DefaultStringSize:         256,                                                                                   // default size for string fields
		DisableDatetimePrecision:  true,                                                                                  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                                                                                  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                                  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                                 // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db1.AutoMigrate(&model.WordInfo{})
	db = db1
}

func SaveWord(wordinfo *model.WordInfo) error {
	// if exists then exit
	if _, b := FindFirst(wordinfo.Word); b == true {
		return nil
	}

	// new create new word info
	ret := db.Create(wordinfo)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

// FindFirst /**
func FindFirst(word string) (*model.WordInfo, bool) {
	var wordinfo model.WordInfo
	ret := db.Where("word = ?", word).First(&wordinfo)
	if ret.Error != nil {
		return nil, false
	}
	return &wordinfo, true
}

/*
FindAll: get All word Info
*/
func FindAll() []model.WordInfo {
	var wordinfos []model.WordInfo
	ret := db.Order("id desc").Find(&wordinfos)
	if ret.Error != nil {
		return nil
	}
	return wordinfos
}
