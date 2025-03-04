package main

import (
	"context"
	"fmt"
	"shop_srvs/inventory_srv/proto"
)

func TestBrandList() {

	BrandListResponse, err := goodsClient.BrandList(context.Background(), &proto.BrandFilterRequest{Pages: 2, PagePerNums: 5})
	if err != nil {
		panic(err)
	}

	for _, brandResponse := range BrandListResponse.Data {
		fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	}
}
