package service

import (
	"context"
	"deviceservice/internal/conf"
	"deviceservice/internal/model"
	"deviceservice/internal/model/messageBus"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

//非凡是的设备接入服务
//1、将非凡是的服务进行封装
//2、将步骤1中的封装函数与边缘网关的物模型服务进行映射
//3、云端通过边缘网关的物模型服务与非凡是打通

var (
	ffsDeviceServiceName = "ffsDeviceService"
	qos                  = 1
)

type ffsDeviceService struct {
	Name             string
	ProductKey       string
	DeviceName       string
	Addr             string
	messageBusClient messageBus.MessageBusClient
	property         []string
	event            []string
	service          []string
	recvChan         chan model.ThingsModel
	sendChan         chan model.ThingsModel
}

func ffsDeviceServerRegister(ffs conf.Ffs) {
	messagebusClient := messageBus.NewMessageBusClient("nanomq")

	mServer := &ffsDeviceService{Name: ffsDeviceServiceName,
		messageBusClient: messagebusClient,
		ProductKey: ffs.ProductKey,
		DeviceName: ffs.DeviceName,
		sendChan : make(chan model.ThingsModel, 1024),
		recvChan : make(chan model.ThingsModel, 1024),
	}
	ss = append(ss, mServer)
}

func (ffs *ffsDeviceService) Start() {
	ctx := context.Background()
	topic := fmt.Sprintf(model.Service, ffs.ProductKey, ffs.DeviceName, "#")
	ffs.messageBusClient.Receive(ctx, topic, ffs.recvChan)
	//todo 接收ffs服务
	go func() {
		//todo 从messagebus中读取消息，发送到nanomq中
		for {
			timer := time.NewTimer(30 * time.Second)
			select {
			case tm := <-ffs.recvChan:
				log.Debug("recvChan:",tm)
				dispatcher(tm)
				timer.Stop()
			case <-timer.C:
				log.Info("ffs messagebus receive timeout after 30 seconds")
				timer.Stop()
			}

		}
	}()

	//todo 定时主动上报ffs属性
	go func() {
		timer := time.NewTimer(10*time.Second)
		for  {
			select {
			case <-timer.C:
				//todo 定时任务
			}
		}
	}()
}

func (mserver *ffsDeviceService) Stop() {
	//todo mqtt device server stop
}
func (mserver *ffsDeviceService) GetName() string {
	return mserver.Name
}

/*
 * func：物模型服务与func匹配
 * 优化：暂时通过switch匹配，后期优化此处动态调用
 */

func dispatcher(tm model.ThingsModel) {
	switch tm.Method {
	case "thing.service.ChangeModel":
		changeModelExample()
	case "thing.service.ChangeBrightness":
		changeBrightnessExample()
	}
}

func changeModelExample() {
	log.Debug("func call changeModelExample!")
	return
}

func changeBrightnessExample() {
	log.Debug("func call changeBrightnessExample!")
	return
}
