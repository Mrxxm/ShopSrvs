package main

import (
	"flag"
	"fmt"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"shop_srvs/order_srv/global"
	"shop_srvs/order_srv/handler"
	"shop_srvs/order_srv/initialize"
	"shop_srvs/order_srv/proto"
	"shop_srvs/order_srv/utils"
	"shop_srvs/order_srv/utils/register/consul"
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
	initialize.InitSrvConn()
	// 3.实例化一个server
	server := grpc.NewServer()
	// 4.注册处理逻辑handler(RegisterGreeterServer为自动生成)
	//proto.RegisterGoodsServer(server, &handler.GoodsService{})
	//proto.RegisterInventoryServer(server, &proto.UnimplementedInventoryServer{})
	//proto.RegisterOrderServer(server, &proto.UnimplementedOrderServer{})
	proto.RegisterOrderServer(server, &handler.OrderServer{})

	// 5.启动服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	// 6.注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	// 7.启动服务
	go func() {
		_ = server.Serve(listener) // 阻塞的方法,需要放在goroutine中，否则后续代码无法执行
	}()
	// 8.服务注册
	// 8.1.初始化配置
	serviceID := fmt.Sprintf("%s", uuid.NewV4()) // 服务id
	register_client := consul.NewRegistryClient(global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	_ = register_client.Register(*IP, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceID)

	// 9.优雅退出，接收终止信号
	quit := make(chan os.Signal) // 无缓冲区通道
	// SIGINT（通常是用户按下 Ctrl+C）和 SIGTERM（通常是终止进程的信号）。当程序收到这两个信号之一时，操作系统会将信号发送到 quit 通道。
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 它用于接收来自操作系统的信号（比如中断信号 SIGINT 或终止信号 SIGTERM）
	<-quit                                               // 这一行代码会阻塞程序的执行，直到从 quit 通道接收到信号
	if err := register_client.DeRegister(serviceID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
