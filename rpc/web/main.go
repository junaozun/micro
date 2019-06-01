package main

import (
        "github.com/micro/go-log"
	    "net/http"

        "github.com/micro/go-web"
        "micro/rpc/web/handler"
)

func main() {
	// create new web service
        service := web.NewService(
                //服务名称
                web.Name("go.micro.web.web"),
                //版本号
                web.Version("latest"),
                //监听端口
                web.Address(":8080"),
        )

	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/web/call", handler.WebCall)

	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
