package service

import (
	"deviceservice/internal/conf"
	"deviceservice/internal/model"
	"deviceservice/internal/model/messageBus"
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var testcase = struct {
	ffs ffsDeviceService

} {
	ffs: ffsDeviceService{
		messageBusClient: messageBus.NewMessageBusClient("nanomq"),
	},
}

func TestChangeModel(t *testing.T) {
	Convey("TestChangeModel", t, func() {
		nanomq := messageBus.NewMessageBusByNanoMQ()
		conf.GetConf()
		Convey("publist message", func() {
			topic := fmt.Sprintf(model.Service, testcase.ffs.ProductKey, testcase.ffs.DeviceName, "ChangeModel")
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
			for {
				select {
				case <-time.After(10 * time.Second):
					t.Log("publish message:", tmjson)
					token := nanomq.Publish(topic, 1, true, tmjson)
					token.Wait()
				}
			}
		})
	})
}
