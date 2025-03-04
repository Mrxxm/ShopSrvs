package main

import (
	"context"
	"fmt"
	"shop_srvs/inventory_srv/proto"
)

func TestGetGoodsList() {

	GoodsListResponse, err := goodsClient.GoodsList(context.Background(), &proto.GoodsFilterRequest{
		TopCategory: 130361,
		KeyWords:    "深海速冻",
	})
	if err != nil {
		panic(err)
	}
	//
	fmt.Println(GoodsListResponse)
	//for _, brandResponse := range AllCategoryResponse.Data {
	//	fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	//}
}

func TestBatchGetGoodsList() {

	GoodsListResponse, err := goodsClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: []int32{428, 429},
	})
	if err != nil {
		panic(err)
	}
	//
	fmt.Println(GoodsListResponse)
	//for _, brandResponse := range AllCategoryResponse.Data {
	//	fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	//}
}
