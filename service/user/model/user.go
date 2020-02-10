package model

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// 完成User类型定义
type User struct {
	Id           int64 `orm:"pk"` // 设置为主键，字段Id, Password首字母必须大写
	Login_name   string
	Pwd          string
	Name         string
	Phone_number string
	Job_number   string
	Permission   int
	Gender       int
	Deleted      bool
}

func init() {
	orm.RegisterModel(new(User)) // 注册模型，建立User类型对象，注册模型时，需要引入包
}

func GetUserListByOffstAndLimit(offset int, limit int) []User {
	var u []User
	o := orm.NewOrm()
	offset = offset - 1
	o.QueryTable("user").Filter("deleted", 0).Offset(offset * limit).Limit(limit).All(&u, "id", "login_name", "name", "phone_number", "job_number", "gender")
	return u
}

func UpdateUserById(change_id int64, user_name string, login_name string, user_phone_number string, user_job_number string, user_gender int) (string, int64) {
	u := User{Id: change_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		u.Name = user_name
		u.Login_name = login_name
		u.Phone_number = user_phone_number
		u.Job_number = user_job_number
		u.Gender = user_gender
		if num, err := o.Update(&u); err == nil {
			logs.Debug("UpdateUserById's num:", num)
			return "success", u.Id
		} else if err != nil {
			logs.Error("UpdateUserById's err:", err)
			return "update faild", -1
		}
	}
	return " user doesn't exist ", -1
}

func AddUser(user_name string, login_name string, user_phone_number string, user_job_number string, user_gender int) (string, int64) {
	//login_name不能重复
	var u User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("login_name", login_name).All(&u)
	if (u != User{}) {
		return "login name have exit,please change another", -1
	}

	//user_gender只能是0或1
	if (user_gender > 1 || user_gender < 0) {
		user_gender = 0
	}

	u.Name = user_name
	u.Login_name = login_name
	u.Phone_number = user_phone_number
	u.Job_number = user_job_number
	u.Gender = user_gender
	id, err := o.Insert(&u)
	if err == nil {
		logs.Debug("AddUser's id:", id)
		return "success", u.Id
	} else {
		logs.Error("AddUser's err:", err)
		return "insert faild", -1
	}
}

func DeleteUserById(delete_id int64) (string, int64) {
	u := User{Id: delete_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		u.Deleted = true
		if num, err := o.Update(&u); err == nil {
			logs.Debug("DeleteUserById's num:", num)
			return "success", u.Id
		} else if err != nil {
			logs.Error("DeleteUserById's err:", err)
			return "delete faild", -1
		}
	}
	return " user doesn't exist ", -1
}

func GetUserById(id int64) (User) {
	u := User{Id: id}
	o := orm.NewOrm()
	err := o.Read(&u, "Id")
	if err != nil {
		return User{}
	} else {
		return u
	}
}

func Login(username string, password string) (*User, bool) {
	u := User{Login_name: username, Pwd: password}
	o := orm.NewOrm()
	err := o.Read(&u, "Login_name", "Pwd")
	if err != nil {
		return nil, false
	} else {
		return &u, true
	}
}

func UpdateUserPwd(user_id int64, old_pwd string, pwd string) string {
	u := User{Id: user_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		if (u.Pwd == old_pwd) {
			u.Pwd = pwd
			if num, err := o.Update(&u, "Pwd"); err == nil {
				logs.Debug("UpdateUserPwd's num:", num)
				return "success"
			} else if err != nil {
				logs.Error("UpdateUserPwd's err:", err)
				return "update faild"
			}
		} else {
			return "old password error"
		}
	}
	return "faild"
}
