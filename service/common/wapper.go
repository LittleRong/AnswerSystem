package common

import (
	"context"
	"errors"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/server"
	"github.com/spf13/viper"
)

// 实现server.HandlerWrapper接口
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logs.Debug("server request", time.Now().Format("2006/1/2 15:04:05"),req.Endpoint())
		logs.Debug("server header", req.Header())
		return fn(ctx, req, rsp)
	}
}

func authWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		logs.Debug("[authWrapper]", req.Endpoint())
		//登陆不需要验证
		if(req.Endpoint()=="UserManage.Login"){
			return fn(ctx, req, rsp)
		}
		header := req.Header()
		if(header==nil){
			//没有则返回错误
			logs.Error("[JWT auth]","get header wrong")
			return errors.New("get header wrong")
		}

		tokenString := header["Authorization"]
		secret := viper.GetString("jwt.secret")

		if(tokenString==""){
			logs.Error("[JWT auth]","no auth meta-data Authorization found in request")
			return errors.New("no auth meta-data Authorization found in request")
		}

		//token校验
		token, err := jwt.Parse(tokenString, secretFunc(secret))

		//失败
		if err != nil {
			logs.Error("[JWT auth]",err)
			return err
		//token校验成功
		} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//ctx.ID = uint64(claims["id"].(float64))
			//ctx.Username = claims["username"].(string)
			logs.Info("ID=", int64(claims["id"].(float64))," username=",claims["username"].(string))
			return fn(ctx, req, rsp)
		} else {
			logs.Error("[JWT auth]",err)
			return err
		}

		return fn(ctx, req, rsp)
	}
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}