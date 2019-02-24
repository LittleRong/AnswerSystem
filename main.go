package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "hello/models"
	_ "hello/routers"
)

var FilterUser = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("user_id").(int)
	beego.Info(v, "+", ok)
	if (ctx.Request.RequestURI != "/check" && ctx.Request.RequestURI != "/index") && !ok {
		ctx.Redirect(302, "/index")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
