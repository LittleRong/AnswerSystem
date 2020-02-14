package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/spf13/pflag"

	_ "web/models"
	_ "web/routers"
	"web/conf"
)

var FilterUser = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("user_id").(int64)
	beego.Info("user_id=",v, "++", ok)
	//if (ctx.Request.RequestURI != "/check" && ctx.Request.RequestURI != "/index") && !ok {
	//	ctx.Redirect(302, "/index")
	//}
}

var FilterPermission = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("permission").(int32)
	beego.Info("permission=",v, "++", ok)
	if (ok && v < 1){
		ctx.Redirect(302, "/404")
	}
}


var (
	cfg = pflag.StringP("config","c","","web config file path")
)

func main() {
	pflag.Parse()
	if err:= conf.Init(*cfg);err != nil {
		panic(err)
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/manage/*", beego.BeforeRouter, FilterPermission)
	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
