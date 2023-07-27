package service

import (
	"deviceservice/internal/model/messageBus"
)


//朗狮的设备接入服务
//1、朗狮是的服务进行封装
//2、将步骤1中的封装函数与边缘网关的物模型服务进行映射
//3、云端通过边缘网关的物模型服务与朗狮打通


var (
	lsDeviceServiceName = "lsDeviceService"
	mqttAddr              = "127.0.0.1:1883"
	lsClientID              = lsDeviceServiceName
)



type landleafDeviceService struct {
	Name string
	messageBusClient messageBus.MessageBusClient
	downlinkTopic string
	uplinkTopic string
}

func landleafDeviceServerRegister() {

}

func (mdds *landleafDeviceService) Start() {

	}


func (mserver *landleafDeviceService) Stop() {
	//todo mqtt device server stop
}
func (mserver *landleafDeviceService) GetName() string {
	return mserver.Name
}

