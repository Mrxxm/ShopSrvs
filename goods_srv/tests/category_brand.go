package main

import (
	"context"
	"fmt"
	"shop_srvs/goods_srv/proto"
)

func TestCategoryBrandList() {

	CategoryBrandResponse, err := goodsClient.CategoryBrandList(context.Background(), &proto.CategoryBrandFilterRequest{
		Pages:       1,
		PagePerNums: 5,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(CategoryBrandResponse)
	//for _, brandResponse := range AllCategoryResponse.Data {
	//	fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	//}
}
