package main

import (
	"github.com/astaxie/beego"
	"service/user/model"
	"context"
	micro "github.com/micro/go-micro"
	proto "service/protoc" //proto文件放置路径
)

type UserManage struct{}

func (this *UserManage) GetUserListByOffstAndLimit(ctx context.Context, req *proto.GetUserListReq, rsp *proto.UserListRsp) error{
	offset := 0
	limit :=10
	userList := model.GetUserListByOffstAndLimit(offset,limit)
	beego.Info("========GetUserListByOffstAndLimit===========",userList)

	return nil
}

func main(){
	//create service
	service := micro.NewService(micro.Name("UserManage"))

	//init
	service.Init()

	//register handler
	proto.RegisterUserManageHandler(service.Server(), new(UserManage))

	//run the server
	if err:=service.Run();err != nil {
		beego.Info("========GetUserListByOffstAndLimit's err===========",err)
	}


}