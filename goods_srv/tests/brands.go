package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/goods_srv/proto"
)

var goodsClient proto.GoodsClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:56311", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	goodsClient = proto.NewGoodsClient(connect)
	if goodsClient == nil {
		panic("创建 gRPC 客户端失败，goodsClient 为 nil")
	}

}

func TestBrandList() {

	BrandListResponse, err := goodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{Pages: 2, PagePerNums: 5})
	if err != nil {
		panic(err)
	}

	for _, brandResponse := range BrandListResponse.Data {
		fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	}
}

func main() {
	Init()
	defer connect.Close()

	TestBrandList()
}
