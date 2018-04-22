package config

import (
	_ "github.com/astaxie/beego/config/xml"
	"github.com/astaxie/beego/config"
	"quickstart/util/log"
)

var consoleLog config.Configer;

func init(){
	iniconf, err := config.NewConfig("ini", "app.conf")
	if err != nil {
		log.Error(err.Error())
	}
	consoleLog = iniconf


}
