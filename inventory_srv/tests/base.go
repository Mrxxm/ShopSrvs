package main

import (
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/inventory_srv/proto"
	"sync"
)

var inventoryClient proto.InventoryClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:59117", grpc.WithInsecure())
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

	//for i := 421; i < 840; i++ {
	//TestSetInventory(int32(i), 100)
	//}
	//
	//TestGetInventory()

	// 主要用于goroutine的执行等待
	var wg sync.WaitGroup
	// 监控goroutine的个数，到执行结束
	wg.Add(80)

	for i := 0; i < 80; i++ {
		go func(i int) {
			defer wg.Done() // wg.Add和wg.Done需要成对出现，否则报错
			TestSell()
			fmt.Println(i)
		}(i)
	}
	// 直到所有协程执行完成，再结束主进程
	wg.Wait()

	//TestGetInventory()
}
