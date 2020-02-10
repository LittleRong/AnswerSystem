package main

import (
	"context"
	"github.com/astaxie/beego/logs"

	"service/common"
	proto "service/protoc/userManage"
	"service/user/model"
)

type UserManage struct{}

func (this *UserManage) GetUserListByOffstAndLimit(ctx context.Context, req *proto.GetUserListReq, rsp *proto.UserListRsp) error {
	offset := 0
	limit := 10
	userList := model.GetUserListByOffstAndLimit(offset, limit)
	//类型转换
	var userMessage []*proto.UserMesssage
	for _, v := range userList {
		u := proto.UserMesssage{Id: int64(v.Id), LoginName: v.Login_name, Name: v.Name, JobNumber: v.Job_number, PhoneNumber: v.Phone_number, Permission: int32(v.Permission), Deleted: v.Deleted, Gender: int32(v.Gender)}
		userMessage = append(userMessage, &u)
	}
	rsp.UserList = userMessage

	return nil
}

func (this *UserManage) UpdateUserById(ctx context.Context, req *proto.ChangeUserReq, rsp *proto.ChangeUserRsp) error {
	var changeId = req.ChangeId
	var name = req.Name
	var loginName = req.LoginName
	var phoneNumber = req.PhoneNumber
	var jobNumber = req.JobNumber
	var gender = req.Gender

	result, id := model.UpdateUserById(changeId, name, loginName, phoneNumber, jobNumber, int(gender))

	rsp.Message = result
	rsp.UserId = id

	return nil
}

func (this *UserManage) AddUser(ctx context.Context, req *proto.AddUserReq, rsp *proto.AddUserRsp) error {
	var name = req.Name
	var loginName = req.LoginName
	var phoneNumber = req.PhoneNumber
	var jobNumber = req.JobNumber
	var gender = req.Gender

	result, id := model.AddUser(name, loginName, phoneNumber, jobNumber, int(gender))

	rsp.Message = result
	rsp.UserId = int64(id)

	return nil
}

func (this *UserManage) DeleteUserById(ctx context.Context, req *proto.DeleteUserReq, rsp *proto.DeleteUserRsp) error {
	var deleteId = req.DeleteId

	result, id := model.DeleteUserById(deleteId)

	rsp.Message = result
	rsp.UserId = id

	return nil
}

func (this *UserManage) GetUserById(ctx context.Context, req *proto.GetUserByIdReq, rsp *proto.UserMesssage) error {
	var userId = req.UserId
	v := model.GetUserById(userId)

	//类型转换
	rsp.Id = v.Id
	rsp.LoginName = v.Login_name
	rsp.Name = v.Name
	rsp.JobNumber = v.Job_number
	rsp.PhoneNumber = v.Phone_number
	rsp.Permission = int32(v.Permission)
	rsp.Deleted = v.Deleted
	rsp.Gender = int32(v.Gender)

	return nil
}

func (this *UserManage) UpdateUserPwd(ctx context.Context, req *proto.UpdatePwdReq, rsp *proto.UpdatePwdRsp) error {
	var userId = req.UserId
	var oldPwd = req.OldPwd
	var newPwd = req.NewPwd

	result := model.UpdateUserPwd(userId, oldPwd, newPwd)

	rsp.Message = result

	return nil
}

func (this *UserManage) Login(ctx context.Context, req *proto.LoginReq, rsp *proto.LoginRsp) error {
	var userName = req.Username
	var pwd = req.Pwd
	user, flag := model.Login(userName, pwd)

	//类型转换
	rsp.UserId = user.Id
	rsp.LoginFlag = flag
	rsp.Permission = int32(user.Permission)

	return nil
}

func main() {
	service,err := common.Init("UserManage")
	if err != nil {
		panic(err)
	}

	//注册服务
	proto.RegisterUserManageHandler(service.Server(), new(UserManage))

	//运行
	if err := service.Run(); err != nil {
		logs.Error("failed-to-do-somthing", err)
	}
}
