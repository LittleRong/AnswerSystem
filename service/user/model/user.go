package model

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id           int `orm:"pk"` // 设置为主键，字段Id, Password首字母必须大写
	Login_name   string
	Pwd          string
	Name         string
	Phone_number string
	Job_number   string
	Permission   int
	Gender       int
	Deleted      bool
}

func GetUserListByOffstAndLimit(offset int, limit int) []User {
	var u []User
	o := orm.NewOrm()
	offset = offset - 1
	o.QueryTable("user").Filter("deleted", 0).Offset(offset * limit).Limit(limit).All(&u, "id", "login_name", "name", "phone_number", "job_number", "gender")
	beego.Info("======GetUserListByOffstAndLimit=====", u)
	return u
}