package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/user_srv/proto"
)

var userClient proto.UserClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	userClient = proto.NewUserClient(connect)
}

func TestGetUserList() {

	UserListResponse, err := userClient.GetUserList(context.Background(), &proto.PageInfo{Page: 1, PageSize: 5})
	if err != nil {
		panic("调用失败")
	}

	for _, userInfoResponse := range UserListResponse.Data {
		fmt.Println(userInfoResponse)
		fmt.Println(userInfoResponse.Id, userInfoResponse.Mobile, userInfoResponse.Nickname)

		checkRes, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: userInfoResponse.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRes.Success)
	}
}

func TestGetUserByMobile() {

	r, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "13777891955"})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.ProtoMessage)
}

func TestGetUserById() {

	r, err := userClient.GetUserList(context.Background(), &proto.PageInfo{Page: 1, PageSize: 5})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(r.ProtoMessage)
}

func main() {
	Init()
	defer connect.Close()

	TestGetUserList()

}
