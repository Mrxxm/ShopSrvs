package main

import (
	"context"
	"fmt"
	"shop_srvs/goods_srv/proto"
)

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
