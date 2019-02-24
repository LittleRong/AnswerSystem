package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "hello/models"
	_ "hello/routers"
)

var FilterUser = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("user_id").(int)
	beego.Info("user_id=",v, "++", ok)
	if (ctx.Request.RequestURI != "/check" && ctx.Request.RequestURI != "/index") && !ok {
		ctx.Redirect(302, "/index")
	}
}

var FilterPermission = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("permission").(int)
	beego.Info("permission=",v, "++", ok)
	if (ok && v < 1){
		ctx.Redirect(302, "/404")
	}
}

func main() {
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/manage/*", beego.BeforeRouter, FilterPermission)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
