package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/goods_srv/proto"
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

	user, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{Mobile: "13777891955"})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(user.Mobile)
}

func TestGetUserById() {

	user, err := userClient.GetUserById(context.Background(), &proto.IdRequest{Id: "3"})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(user.Id, user.Mobile)
}

func TestCreateUser() {

	user, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Nickname: "xxmx",
		Mobile:   "13777891966",
		Password: "admin321",
	})
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(user.Nickname, user.Mobile)
}

func main() {
	Init()
	defer connect.Close()

	TestGetUserList()
	TestGetUserByMobile()
	TestGetUserById()
	TestCreateUser()
}
