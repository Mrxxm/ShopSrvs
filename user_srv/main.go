package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"shop_srvs/user_srv/global"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	// 6.注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	// 7.服务注册
	// 7.1.初始化配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)

	// 7.2.创建一个consul客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 7.3.注册服务&生成注册对象&生成检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.15.21:50051"),
		Timeout:                        "5s",  // 超时时间
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "10s", // 多久后注销服务
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	registration.ID = global.ServerConfig.Name
	registration.Port = *Port
	registration.Tags = []string{"xxm", "grpc", "user", "srv"}
	registration.Address = "192.168.15.21"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	_ = server.Serve(listener)

}
