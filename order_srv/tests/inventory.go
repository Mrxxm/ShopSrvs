package main

import (
	"context"
	"fmt"
	"shop_srvs/order_srv/proto"
)

func TestSetInventory(goodsId, Num int32) {

	_, _ = inventoryClient.SetInv(context.Background(), &proto.GoodsInvInfo{
		GoodsId: goodsId,
		Num:     Num,
	})

	fmt.Println("设置成功")
}

func TestGetInventory() {

	GoodsInvInfo, err := inventoryClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: 421,
		Num:     100,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(GoodsInvInfo)
}

func TestSell() {
	GoodsInvInfo, err := inventoryClient.Sell(context.Background(), &proto.SellInfo{
		GoodsInfo: []*proto.GoodsInvInfo{
			{
				GoodsId: 421,
				Num:     1,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(GoodsInvInfo)
}
