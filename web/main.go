package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/spf13/pflag"

	"web/conf"
	_ "web/models"
	_ "web/routers"
)

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("user_id").(int64)
	beego.Info("00000RequestURI0000000", ctx.Request.RequestURI, " ", ctx.Input, "  body=", string(ctx.Input.RequestBody))
	if (ctx.Request.RequestURI != "/v1/index/check" && ctx.Request.RequestURI != "/v1/index") && !ok {
		ctx.Redirect(302, "v1/index")
	}
}

var FilterPermission = func(ctx *context.Context) {
	v, ok := ctx.Input.Session("permission").(int32)
	if (ok && v < 1) {
		ctx.Redirect(302, "/404")
	}
}

var (
	cfg = pflag.StringP("config", "c", "", "web config file path")
)

func main() {
	pflag.Parse()
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/manage/*", beego.BeforeRouter, FilterPermission)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.SetStaticPath("/swagger", "swagger")

	//打开session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
