package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"shop_srvs/user_srv/global"
	"shop_srvs/user_srv/handler"
	"shop_srvs/user_srv/initialize"
	"shop_srvs/user_srv/proto"
	"shop_srvs/user_srv/utils"
	"syscall"
)

func main() {
	// 使用命令行设置IP和端口号
	IP := flag.String("ip", "192.168.15.21", "ip地址")
	Port := flag.Int("port", 0, "端口号")
	flag.Parse()
	// 1.日志初始化
	initialize.InitLogger()
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("ip: , Port: ", *IP, " ", *Port)

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
		GRPC:                           fmt.Sprintf("%s:%d", *IP, *Port),
		Timeout:                        "5s",  // 超时时间
		Interval:                       "5s",  // 健康检查间隔
		DeregisterCriticalServiceAfter: "10s", // 多久后注销服务
	}
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serviceID := fmt.Sprintf("%s", uuid.NewV4()) // 服务id
	//registration.ID = global.ServerConfig.Name
	registration.ID = serviceID // 启动多个服务
	registration.Port = *Port
	registration.Tags = []string{"xxm", "grpc", "user", "srv"}
	registration.Address = "192.168.15.21"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	go func() {
		_ = server.Serve(listener) // 阻塞的方法,需要放在goroutine中，否则后续代码无法执行
	}()

	// 8.优雅退出，接收终止信号
	quit := make(chan os.Signal) // 无缓冲区通道
	// SIGINT（通常是用户按下 Ctrl+C）和 SIGTERM（通常是终止进程的信号）。当程序收到这两个信号之一时，操作系统会将信号发送到 quit 通道。
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 它用于接收来自操作系统的信号（比如中断信号 SIGINT 或终止信号 SIGTERM）
	<-quit                                               // 这一行代码会阻塞程序的执行，直到从 quit 通道接收到信号
	if err = client.Agent().ServiceDeregister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
