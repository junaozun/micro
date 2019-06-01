package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"micro/rpc/srv/handler"
	srv "micro/rpc/srv/proto/srv"
)

func main() {
	// New Service
	service := micro.NewService(
		//服务名
		micro.Name("go.micro.srv.srv"),
		//版本号
		micro.Version("latest"),
	)

	//初始化服务
	service.Init()

	//注册服务
	srv.RegisterSrvHandler(service.Server(), new(handler.Srv))

	// Register Struct as Subscriber
	//注册结构体，源自于Subscriber文件夹
	//micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), new(subscriber.Srv))

	// Register Function as Subscriber
	//注册一个方法，源自于Subscriber文件夹
	//micro.RegisterSubscriber("go.micro.srv.srv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
