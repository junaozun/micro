package handler

import (
	"context"
	"encoding/json"
	web "micro/rpc/srv/proto/srv"
	"net/http"
	"time"
	"github.com/micro/go-micro/client"
)

//（传出，浏览器传入的参数）
func WebCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	//创建一个map
	var request map[string]interface{}
	//获取浏览器传入的参数实体r.body，解码到request的map里面
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

    //调用服务，返回句柄
	webClient := web.NewSrvService("go.micro.srv.srv", client.DefaultClient)
	//通过句柄，调用call函数
	rsp, err := webClient.Call(context.TODO(), &web.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//接受服务调用的返回信息，创建成为map
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	//将response转换成json格式，发送给前端
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}


