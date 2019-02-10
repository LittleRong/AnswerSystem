package user

import (
	"fmt"
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
	err = o.Read(u,"Id","Pwd")
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

func IsManager(ID int) (*User,bool) {
	u := User{Id:ID}
	o := orm.NewOrm()
	err := o.Read(&u,"Id")
	if err != nil {
		return nil,false
	} else if u.Permission == 1 {
		return &u,true
	} else if u.Permission == 2 {
		return &u,true
	} else {
		return &u,false
	}
}

func Login(username string,password string) (*User,bool) {
	u := User{Login_name:username,Pwd:password}
	o := orm.NewOrm()
	err := o.Read(&u,"Login_name","Pwd")
	if err != nil {
		return nil,false
	} else {
		return &u,true
	}
}

func GetUserById(id int) (User) {
	u := User{Id:id}
	o := orm.NewOrm()
	err := o.Read(&u,"Id")
	if err != nil {
		return User{}
	} else {
		return u
	}
}
