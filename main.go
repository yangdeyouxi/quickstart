package main

import (
	_ "log"
	_ "quickstart/routers"
	_ "quickstart/util/goroutine"
	_ "quickstart/module/reptile"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()


}

