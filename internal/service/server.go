package service

import (
	log "github.com/sirupsen/logrus"
	"deviceservice/internal/conf"
	"os"
	"os/signal"
)

type Server interface {
	Start()
	Stop()
	GetName() string
}

var (
	ss[]Server
)

func RegisterService(config conf.Config)  {
	ffsDeviceServerRegister(config.Ffs)
	GatewayDeviceServerRegister(config.Gateway)
}


func Start()  {
	c := make(chan os.Signal)
	signal.Notify(c,os.Interrupt,os.Kill)
	//todo for each registed service,and start
	for k := range ss {
		log.Info("Start server:",ss[k].GetName())
		go ss[k].Start()
	}
	s := <-c
	log.Println("End...", s)
}



