package model

type Conf struct {
	Port     string `yaml:"port"`
	RtimeOut int    `yaml:"rtimeOut"`
	WtimeOut int    `yaml:"wtimeOut"`
}
