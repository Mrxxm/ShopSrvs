package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/order_srv/proto"
)

var orderClient proto.OrderClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:53864", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	orderClient = proto.NewOrderClient(connect)
	if orderClient == nil {
		panic("创建 gRPC 客户端失败，inventoryClient 为 nil")
	}

}

func TestCreateCartItem(userId, nums, goodIds int32) {

	rep, err := orderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  userId,
		Nums:    nums,
		GoodsId: goodIds,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("设置成功", rep.Id)
}

func TestCartItemList(userId int32) {
	rep, err := orderClient.CartItemList(context.Background(), &proto.UserInfo{
		Id: userId,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("设置成功", rep.Data)
}

func TestUpdateCartItem(id, goodId, userId int32) {
	_, err := orderClient.UpdateCartItem(context.Background(), &proto.CartItemRequest{
		Id:      id,
		GoodsId: goodId,
		UserId:  userId,
		Checked: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCreateOrder() {
	_, err := orderClient.Create(context.Background(), &proto.OrderRequest{
		UserId:  1,
		Address: "杭州市",
		Name:    "xxm",
		Mobile:  "13777891966",
		Post:    "请尽快发货",
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	Init()
	defer connect.Close()

	//TestCreateCartItem(1, 2, 421)
	//TestCartItemList(1)
	//TestUpdateCartItem(1, 421, 1)
	TestCreateOrder()
}
