package main

import (
	_ "hello/routers"
	_ "hello/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

