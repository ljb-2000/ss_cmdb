package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "ss_cmdb/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	fmt.Println("main starting...")

	beego.Run()
}
