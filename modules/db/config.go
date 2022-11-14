package db

import (
	"howToUseMod/modules/helpers"
	"log"

	"gopkg.in/ini.v1"
)

type DBConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config DBConfigList

func SetConfig() {
	log.Println("init()が実行されました")
	LoadSQLiteConfig()
	helpers.LogingSettings(Config.LogFile)
}

func LoadSQLiteConfig() {
	config, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = DBConfigList{
		Port:      config.Section("web").Key("port").MustString("8080"),
		SQLDriver: config.Section("db").Key("driver").String(),
		DbName:    config.Section("db").Key("name").String(),
		LogFile:   config.Section("web").Key("logfile").String(),
	}
}
