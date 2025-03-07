package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop_srvs/order_srv/global"
	"shop_srvs/order_srv/proto"
)

func InitSrvConn() {
	goodsConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("连接 [goods_srv] 失败", err.Error())
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodsConn)

	inventoryConn, err := grpc.Dial(fmt.Sprintf("consul://%s:%d/%s?wait=14s", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port, global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("连接 [inventory_srv] 失败", err.Error())
	}

	global.InventorySrvClient = proto.NewInventoryClient(inventoryConn)
}
