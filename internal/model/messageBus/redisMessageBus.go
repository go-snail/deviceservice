package messageBus

import (
	"context"
	"deviceservice/internal/conf"
	"deviceservice/internal/model"
	"deviceservice/internal/utils"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var (
	readChannelLen = 1024
	sendChannellen = 1024
)

type redisMessageBus struct {
	ProductKey string
	DeviceName string
	rc          *redis.Client
	readChannel chan model.ThingsModel
	sendChannel chan model.ThingsModel
}

func NewMessageBusByRedis() *redisMessageBus {
	return &redisMessageBus{
		rc:          utils.NewRedisClient(),
		readChannel: make(chan model.ThingsModel, readChannelLen),
		sendChannel: make(chan model.ThingsModel, sendChannellen),
	}
}
func (rmb *redisMessageBus) Receive(ctx context.Context,topic string,reChan chan model.ThingsModel) {
	setKey := fmt.Sprintf(model.Set, conf.C.Gateway.ProductKey, conf.C.Gateway.DeviceName)
	res, err := rmb.rc.XReadGroup(ctx, &redis.XReadGroupArgs{
		// Streams第二个参数为ID，list of streams and ids, e.g. stream1 stream2 id1 id2
		// id为 >，表示最新未读消息ID，也是未被分配给其他消费者的最新消息
		// id为 0 或其他，表示可以获取已读但未确认的消息。这种情况下BLOCK和NOACK都会忽略
		// id为具体ID，表示获取这个消费者组的pending的历史消息，而不是新消息
		Streams:  []string{setKey, ">"},
		Group:    "cg1", //消费者组名
		Consumer: "c1",  // 消费者名
		Count:    1,
		Block:    0,    // 是否阻塞，=0 表示阻塞且没有超时限制。只要大于1条消息就立即返回
		NoAck:    true, // true-表示读取消息时确认消息
	}).Result()
	if err != nil {
		log.Errorln("WriteMessage:", err.Error())
	}
	id := res[0].Messages[0].ID
	var payload = res[0].Messages[0].Values[id]
	if p, ok := payload.(model.ThingsModel); ok {
		rmb.readChannel <- p
		rmb.rc.XDel(ctx, setKey, id)
		return
	}
	log.Errorln("Message Get Error message type!!")
	rmb.rc.XDel(ctx, setKey, id)
	return
}
func (rmb *redisMessageBus) Send(ctx context.Context,topic string, sendChan chan model.ThingsModel) {
	var tm model.ThingsModel
	postKey := fmt.Sprintf(model.Post, rmb.ProductKey,rmb.DeviceName)
	err := rmb.rc.XAdd(ctx, &redis.XAddArgs{
		Stream:     postKey, // 设置流stream的 key，消息队列名
		NoMkStream: false,   //为false，key不存在会新建
		MaxLen:     10000,   //消息队列最大长度，队列长度超过设置最大长度后，旧消息会被删除
		Approx:     false,   //默认false，设为true时，模糊指定stram的长度
		ID:         "*",     //消息ID，* 表示由Redis自动生成
		Values:     tm,
		// MinID: "id",//超过设置长度值，丢弃小于MinID消息id
		// Limit: 1000, //限制长度，基本不用
	}).Err()
	if err != nil {
		log.Errorln("WriteMessage:", err.Error())
	}
}

//func (rmb *redisMessageBus)PostReply(ctx context.Context,tm *model.ThingsModel) {
//
//}

//
//func (rmb *redisMessageBus)SetReply(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *redisMessageBus)Event(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *redisMessageBus)EventReply(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *redisMessageBus)Service(ctx context.Context,tm *model.ThingsModel){
//
//}
//func (rmb *redisMessageBus)ServiceReply(ctx context.Context,tm *model.ThingsModel){
//
//}
