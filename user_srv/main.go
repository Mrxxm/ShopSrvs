package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"shop_srvs/user_srv/handler"
	"shop_srvs/user_srv/initialize"
	"shop_srvs/user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "端口号")
	flag.Parse()
	// 1.日志初始化
	initialize.InitLogger()
	zap.S().Info("ip: , Port: ", *IP, *Port)

	// 2.初始化配置
	initialize.InitConfig()
	initialize.InitDB()
	// 3.实例化一个server
	server := grpc.NewServer()
	// 4.注册处理逻辑handler(RegisterGreeterServer为自动生成)
	proto.RegisterUserServer(server, &handler.UserService{})
	// 5.启动服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	_ = server.Serve(listener)
}
