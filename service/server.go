package service

import (
	log "github.com/sirupsen/logrus"
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

func init()  {
	HttpServerRegister()
}


func Start()  {
	c := make(chan os.Signal)
	signal.Notify(c)
	//todo for each registed service,and start
	for k := range ss {
		log.Info("Start server:",ss[k].GetName())
		go ss[k].Start()
	}
	s := <-c
	log.Println("End...", s)
}



