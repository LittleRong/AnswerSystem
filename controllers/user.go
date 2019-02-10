package controllers

import (
	"hello/models/user"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation" // 用于校验信息
)

type UserController struct {
	beego.Controller
}

func (c *UserController) PageLogin() {
	c.TplName = "login.html" // 将hello.html页面输出
}

func (c *UserController) Register() {
	id, _ := c.GetInt("userid") // login.html中传过来的数据，这个GetInt返回两个值
	password := c.GetString("password")
	fmt.Println("This is id and password")
	fmt.Println(id, password)

	valid := validation.Validation{}
	valid.Required(id, "id") // 校验是否为空值
	valid.Required(password, "password")
	// valid.MaxSize(id, 20, "id")
	switch { // 使用switch方式来判断是否出现错误，如果有错，则打印错误并返回
	case valid.HasErrors():
		fmt.Println(valid.Errors[0].Key + valid.Errors[0].Message)
		c.TplName = "bad.html"
		return
	}

	u := &user.User{
		Id:       id,
		Pwd: password,
	}

	err := u.Create()
	if err != nil {
		fmt.Println(err)
		c.TplName = "bad.html"
		return
	}
	c.TplName = "welcome.html"
}

func (c *UserController) Reallogin() {
	id, _ := c.GetInt("userid")
	password := c.GetString("password")
	u := &user.User{
		Id:       id,
		Pwd: password,
	}
	err := u.ReadDB()
	if err != nil {
		fmt.Println(err)
		c.TplName = "bad.html"
		return
	}
	c.TplName = "welcome.html"
}

func Index(){


}