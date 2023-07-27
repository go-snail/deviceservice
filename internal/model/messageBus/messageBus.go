package messageBus

import (
	"context"
	"deviceservice/internal/model"
)

type MessageBusClient interface {
	 Receive(ctx context.Context,topic string,recvChan chan model.ThingsModel)
	// PostReply(ctx context.Context,tm *model.ThingsModel)
	 Send(ctx context.Context,topic string,sendChan chan model.ThingsModel)
	 //SetReply(ctx context.Context,tm *model.ThingsModel)
	 //Event(ctx context.Context,tm *model.ThingsModel)
	 //EventReply(ctx context.Context,tm *model.ThingsModel)
	// Service(ctx context.Context,tm *model.ThingsModel)
	 //ServiceReply(ctx context.Context,tm *model.ThingsModel)
}

type Message struct {

}

func NewMessageBusClient(t string) (MessageBusClient) {
	if t == "redis" {
		return NewMessageBusByRedis()
	} else {
		return NewMessageBusByNanoMQ()
	}
}
//
//func WriteMessage(key string, tm model.ThingsModel) {
//
//	rconn := GetRedisConnect()
//	err := rconn.XAdd(ctx,&redis.XAddArgs{
//		Stream:     key, // 设置流stream的 key，消息队列名
//		NoMkStream: false,         //为false，key不存在会新建
//		MaxLen:     10000,         //消息队列最大长度，队列长度超过设置最大长度后，旧消息会被删除
//		Approx:     false,         //默认false，设为true时，模糊指定stram的长度
//		ID:         "*",           //消息ID，* 表示由Redis自动生成
//		Values: tm,
//		// MinID: "id",//超过设置长度值，丢弃小于MinID消息id
//		// Limit: 1000, //限制长度，基本不用
//	}).Err()
//	if err != nil {
//		log.Errorln("WriteMessage:",err.Error())
//	}
//}



