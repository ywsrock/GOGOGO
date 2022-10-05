package model

import "gorm.io/gorm"

type WordInfo struct {
	gorm.Model
	Word            string // Word
	Syllable        string `gorm:"default:''"` // Syllable
	PhoneticSymbols string `gorm:"default:''"` // Phonetic symbols and reading
	Info            string // Abbreviation Meaning
	InfosJp         string `gorm:"type:text"` // Meaning list
	InfosEn         string `gorm:"type:text"`
	VoiceLink       string // voice link
	Memo            string // memo
}
