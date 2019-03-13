package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"service/user/model"

	_ "github.com/go-sql-driver/mysql"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc/userManage" //proto文件放置路径
)

type UserManage struct{}

func (this *UserManage) GetUserListByOffstAndLimit(ctx context.Context, req *proto.GetUserListReq, rsp *proto.UserListRsp) error{
	offset := 0
	limit :=10
	userList := model.GetUserListByOffstAndLimit(offset,limit)
	beego.Info("========GetUserListByOffstAndLimit000===========",userList)
	//类型转换
	var userMessage []*proto.UserMesssage
	for _,v := range userList {
		u := proto.UserMesssage{Id:int64(v.Id),LoginName:v.Login_name,Name:v.Name,JobNumber:v.Job_number,PhoneNumber:v.Phone_number,Permission:int32(v.Permission),Deleted:v.Deleted,Gender:int32(v.Gender)}
		userMessage = append(userMessage,&u)
	}
	rsp.UserList = userMessage

	return nil
}

func (this *UserManage) UpdateUserById(ctx context.Context, req *proto.ChangeUserReq, rsp *proto.ChangeUserRsp) error{
	var changeId = req.ChangeId
	var name = req.Name
	var loginName = req.LoginName
	var phoneNumber = req.PhoneNumber
	var jobNumber = req.JobNumber
	var gender = req.Gender

	result,id := model.UpdateUserById(changeId,name,loginName,phoneNumber,jobNumber,int(gender))

	rsp.Message = result
	rsp.UserId = id

	return nil
}

func (this *UserManage) AddUser(ctx context.Context, req *proto.AddUserReq, rsp *proto.AddUserRsp) error{
	var name = req.Name
	var loginName = req.LoginName
	var phoneNumber = req.PhoneNumber
	var jobNumber = req.JobNumber
	var gender = req.Gender

	result,id := model.AddUser(name,loginName,phoneNumber,jobNumber,int(gender))

	rsp.Message = result
	rsp.UserId = int64(id)

	return nil
}

func (this *UserManage) DeleteUserById(ctx context.Context, req *proto.DeleteUserReq, rsp *proto.DeleteUserRsp) error{
	var deleteId = req.DeleteId

	result,id := model.DeleteUserById(deleteId)

	rsp.Message = result
	rsp.UserId = id

	return nil
}

func (this *UserManage) GetUserById(ctx context.Context, req *proto.GetUserByIdReq, rsp *proto.UserMesssage) error{
	var userId = req.UserId
	v := model.GetUserById(userId)

	//类型转换
	rsp = &proto.UserMesssage{Id:v.Id,LoginName:v.Login_name,Name:v.Name,JobNumber:v.Job_number,PhoneNumber:v.Phone_number,Permission:int32(v.Permission),Deleted:v.Deleted,Gender:int32(v.Gender)}
	return nil
}

func main(){

	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)

	//create service
	service := micro.NewService(micro.Name("UserManage"))

	//init
	service.Init()

	//register handler
	proto.RegisterUserManageHandler(service.Server(), new(UserManage))

	//run the server
	if err:=service.Run();err != nil {
		beego.Info("========UserManage's err===========",err)
	}
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:ganxiaorong0703@tcp(localhost:3306)/problem?charset=utf8")
}