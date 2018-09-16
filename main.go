package main

import (
	_ "day04/routers"
	"github.com/astaxie/beego"
	_ "day04/models"
	_ "day04/tools"
)

func main() {
	beego.Run()
}

