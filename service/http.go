package service

import (
	"ds-auth/controller"
	"github.com/gin-gonic/gin"
)

var (
	name = "HttpServer"
)
type HttpServer struct {
	Name string
}

func HttpServerRegister() {
	hServer := &HttpServer{Name: name}
	ss = append(ss,hServer)
}


func (hserver *HttpServer)Start()  {
	//todo httpserver start
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/mqtt/auth",controller.Auth)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func (hserver *HttpServer)Stop()  {
	//todo httpserver stop
}
func (hserver *HttpServer)GetName() string {
	return hserver.Name
}


