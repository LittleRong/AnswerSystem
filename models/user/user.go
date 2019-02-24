package user

import (
	"fmt"
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

func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u, "Id", "Pwd")
	return err
}

func (u *User) Create() (err error) {
	o := orm.NewOrm()
	fmt.Println("Create success!")
	_, _ = o.Insert(u)
	return err
}

func (u *User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(u)
	return err
}

func IsManager(ID int) (*User, bool) {
	u := User{Id: ID}
	o := orm.NewOrm()
	err := o.Read(&u, "Id")
	if err != nil {
		return nil, false
	} else if u.Permission == 1 {
		return &u, true
	} else if u.Permission == 2 {
		return &u, true
	} else {
		return &u, false
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

func GetUserById(id int) (User) {
	u := User{Id: id}
	o := orm.NewOrm()
	err := o.Read(&u, "Id")
	if err != nil {
		return User{}
	} else {
		return u
	}
}

func UpdateUserPwd(user_id int, old_pwd string, pwd string) string {
	u := User{Id: user_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		if (u.Pwd == old_pwd) {
			u.Pwd = pwd
			if num, err := o.Update(&u, "Pwd"); err == nil {
				beego.Info("======UpdateUserPwd's num=====", num)
				return "success"
			} else if err != nil {
				beego.Info("======UpdateUserPwd's err=====", err)
				return "update faild"
			}
		} else {
			return "old password error"
		}
	}
	return "faild"
}

func GetUserListByOffstAndLimit(offset int, limit int) []User {
	var u []User
	o := orm.NewOrm()
	offset = offset - 1
	o.QueryTable("user").Filter("deleted", 0).Offset(offset * limit).Limit(limit).All(&u, "id", "login_name", "name", "phone_number", "job_number", "gender")
	beego.Info("======GetUserListByOffstAndLimit=====", u)
	return u
}

func UpdateUserById(change_id int, user_name string, login_name string, user_phone_number string, user_job_number string, user_gender int) string {
	u := User{Id: change_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		u.Name = user_name
		u.Login_name = login_name
		u.Phone_number = user_phone_number
		u.Job_number = user_job_number
		u.Gender = user_gender
		if num, err := o.Update(&u); err == nil {
			beego.Info("======UpdateUserById's num=====", num)
			return "success"
		} else if err != nil {
			beego.Info("======UpdateUserById's err=====", err)
			return "update faild"
		}
	}
	return " user doesn't exist "
}

func AddUser(user_name string, login_name string, user_phone_number string, user_job_number string, user_gender int) string {
	//login_name不能重复
	var u User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("login_name", login_name).All(&u)
	if (u != User{}) {
		return "login name have exit,please change another"
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
		beego.Info("======AddUser's id=====", id)
		return "success"
	} else {
		beego.Info("======AddUser's err=====", err)
		return "insert faild"
	}
}

func DeleteUserById(delete_id int) string {
	u := User{Id: delete_id}
	o := orm.NewOrm()
	if o.Read(&u) == nil {
		u.Deleted = true
		if num, err := o.Update(&u); err == nil {
			beego.Info("======DeleteUserById's num=====", num)
			return "success"
		} else if err != nil {
			beego.Info("======DeleteUserById's err=====", err)
			return "delete faild"
		}
	}
	return " user doesn't exist "
}
