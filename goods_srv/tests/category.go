package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
)

func TestAllCategoryList() {

	AllCategoryResponse, err := goodsClient.GetAllCategorysList(context.Background(), &empty.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Println(AllCategoryResponse)
	//for _, brandResponse := range AllCategoryResponse.Data {
	//	fmt.Println(brandResponse.Id, brandResponse.Name, brandResponse.Logo)
	//}
}
