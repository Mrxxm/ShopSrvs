package main

import (
	"google.golang.org/grpc"
	"shop_srvs/goods_srv/proto"
)

var goodsClient proto.GoodsClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:57562", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	goodsClient = proto.NewGoodsClient(connect)
	if goodsClient == nil {
		panic("创建 gRPC 客户端失败，goodsClient 为 nil")
	}

}

func main() {
	Init()
	defer connect.Close()

	//TestBrandList()
	//TestAllCategoryList()
	//TestCategoryBrandList()
	TestBatchGetGoodsList()
}
