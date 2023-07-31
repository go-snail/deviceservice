package messageBus

import (
	"context"
	"deviceservice/internal/model"
	"deviceservice/internal/utils"
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	addr     = "127.0.0.1:1883"
	clientid = ""
)

type nanoMQMessageBus struct {
	mqtt.Client
}

func NewMessageBusByNanoMQ() *nanoMQMessageBus {
	nmqmb, err := utils.NewNanoMQClient(addr, clientid)
	if err != nil {
		log.Error("NewMessageBusByNanoMQ failed:", err)
		os.Exit(utils.MessageBusClientInitialErr)
	}
	return &nanoMQMessageBus{
		nmqmb,
	}
}

/*
从设备直连的nanomq
*/
func (nanomb *nanoMQMessageBus) Receive(ctx context.Context, topic string, recvChan chan model.ThingsModel) {
	log.Info(topic)
	if token := nanomb.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
		log.Infof("TOPIC: %s\n", message.Topic())
		log.Infof("MSG: %s\n", message.Payload())
		//todo read message form nanomq
		var tm model.ThingsModel
		if err := json.Unmarshal(message.Payload(), &tm); err != nil {
			log.Error("unmarshal err:", err)
			return
		}
		recvChan <- tm
	}); token.Wait() && token.Error() != nil {
		log.Error("subscribe token err:", token.Error())
	}

}

// func (rmb *nanoMQMessageBus)PostReply(ctx context.Context,tm *model.ThingsModel) {
//
// }
func (nanomb *nanoMQMessageBus) Send(ctx context.Context, topic string, tm model.ThingsModel) {
	tmjson, err := json.Marshal(tm)
	if err != nil {
		log.Error("thingsModel marshal err!")
		return
	}
	log.Debug("publish message:", string(tmjson))
	token := nanomb.Publish(topic, 1, true, tmjson)
	token.Wait()

}

//func (rmb *nanoMQMessageBus)SetReply(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *nanoMQMessageBus)Event(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *nanoMQMessageBus)EventReply(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *nanoMQMessageBus)Service(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *nanoMQMessageBus)ServiceReply(ctx context.Context,tm *model.ThingsModel){
//
//}

func (nanomb *nanoMQMessageBus) subscribe() {

}

// 按照deviceservice配置文件，遍历subscribe service服务
func subscribeService() {

}

func subscribeSetProperty() {

}

func subscribePostPropertyReply() {

}
