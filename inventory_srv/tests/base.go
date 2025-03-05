package main

import (
	"google.golang.org/grpc"
	"shop_srvs/inventory_srv/proto"
)

var inventoryClient proto.InventoryClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:57895", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	inventoryClient = proto.NewInventoryClient(connect)
	if inventoryClient == nil {
		panic("创建 gRPC 客户端失败，inventoryClient 为 nil")
	}

}

func main() {
	Init()
	defer connect.Close()

	//TestSetInventory(421, 100)
	TestGetInventory()
	TestSell()
	TestGetInventory()
}
