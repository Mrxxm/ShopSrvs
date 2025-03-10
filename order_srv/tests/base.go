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
	connect, err = grpc.Dial("192.168.15.21:53681", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	orderClient = proto.NewOrderClient(connect)
	if orderClient == nil {
		panic("创建 gRPC 客户端失败，inventoryClient 为 nil")
	}

}

func TestCreateCartItem() {

	rep, err := orderClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  1,
		Nums:    2,
		GoodsId: 421,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("设置成功", rep.Id)
}

func main() {
	Init()
	defer connect.Close()

	TestCreateCartItem()
}
