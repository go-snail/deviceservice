package messageBus

import (
	"deviceservice/internal/conf"
	"deviceservice/internal/model"
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"testing"
	"time"
)

var testcase = struct {
	nanombt MessageBusClient
	productkey string
	devicename string
	service string
} {
	nanombt : NewMessageBusClient("nanomq"),
	productkey: "product",
	devicename: "devicename",
	service: "service",
}

func TestSendThingModel(t *testing.T) {
	Convey("TestSendThingModel", t, func() {

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
			tmjson, err := json.Marshal(tm)
			So(err, ShouldBeNil)
			sendChan := make(chan model.ThingsModel, 1024)
			for {
				select {
				case <-time.After(10 * time.Second):
					t.Log("publish message:", tmjson)
					testcase.nanombt.Send(context.Background(), topic, sendChan)
				}
			}
		})
	})
}

func TestRecvThingModel(t *testing.T) {
	Convey("TestSendThingModel", t, func() {
		ctx := context.Background()
		topic := fmt.Sprintf(model.Service, conf.C.Ffs.ProductKey, conf.C.Ffs.DeviceName, "ChangeModel")
		recvChan := make(chan model.ThingsModel, 1024)
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
