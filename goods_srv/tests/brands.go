package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"shop_srvs/goods_srv/proto"
)

var brandClient proto.GoodsClient
var connect *grpc.ClientConn
var err error

func Init() {
	connect, err = grpc.Dial("192.168.15.21:53296", grpc.WithInsecure())
	if err != nil {
		panic("连接失败")
	}

	brandClient = proto.NewGoodsClient(connect)
}

func TestBrandList() {

	BrandListResponse, err := brandClient.BrandList(context.Background(), &proto.BrandFilterRequest{Pages: 1, PagePerNums: 5})
	if err != nil {
		panic(err.Error())
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
