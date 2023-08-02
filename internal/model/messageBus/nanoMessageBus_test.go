package messageBus

import (
	"context"
	"deviceservice/internal/model"
	"fmt"
	"github.com/rs/xid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)
var testcase = struct {
	nanombt MessageBusClient
	productkey string
	devicename string
	service string
	sendChan chan model.ThingsModel
	recvChan chan model.ThingsModel
} {
	nanombt : NewMessageBusClient("nanomq"),
	productkey: "product",
	devicename: "devicename",
	service: "ChangeModel",

	recvChan: make(chan model.ThingsModel,1024),
}
func TestNanoMQSendMessage(t *testing.T)  {
	Convey("TestNanoMQSendMessage", t, func() {
		Convey("publist message", func() {

			topic := fmt.Sprintf(model.Service, testcase.productkey, testcase.devicename, testcase.service)
			tm := new(model.ThingsModel)
			tm.Id = xid.New().String()
			tm.Version = "v1.0"

			tm.Ack = struct{ Ack bool }{Ack: false}
			param := make(map[string]model.Identifer)
			iden1 := new(model.Identifer)
			iden1.Value = "on"
			iden1.Time = time.Now()

			iden2 := new(model.Identifer)
			iden2.Value = "3"
			iden2.Time = time.Now()

			param["Power"] = *iden1
			param["WF"] = *iden2
			tm.Payload = param

			//sendChan := make(chan model.ThingsModel, 1024)
			for {
				select {
				case <-time.After(10 * time.Second):
					//testcase.nanombt.sendChan <- *tm
					t.Log("topic:",topic)
					testcase.nanombt.Send(context.Background(), topic,*tm)
				}
			}
		})
	})
}


func TestNanoMQRecvMessage(t *testing.T) {
	Convey("TestRecvThingModel", t, func() {
		ctx := context.Background()
		topic := fmt.Sprintf(model.Service, testcase.productkey, testcase.devicename, testcase.service)
		recvChan := make(chan model.ThingsModel, 1024)
		t.Log("recv topic:", topic)
		testcase.nanombt.Receive(ctx, topic, recvChan)

		for {
			timer := time.NewTimer(15 * time.Second)
			select {
			case tm := <-recvChan:
				t.Log("recvChan receive message:", tm)
			case <-timer.C:
				t.Log("recv timeout!")
			}
		}
	})
}