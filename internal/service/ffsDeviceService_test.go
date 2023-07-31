package service

import (
	"deviceservice/internal/model/messageBus"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

	})
}
