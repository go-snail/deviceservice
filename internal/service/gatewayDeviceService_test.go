package service

import (
	."github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestGateway struct {

}

func TestGatewayUplinkCPU(t *testing.T)  {
	Convey("TestGatewayUplinkCPU",t, func() {
		cpu := getCPU()
		for i,_ := range cpu {
			t.Log("cpu percentage:",cpu[i])
		}
	})
}


func TestGatewayUplinkMemory(t *testing.T) {
	Convey("TestGatewayUplinkMemory",t, func() {
		mem := getMemory()
		t.Log("memory:",mem.Total)
		t.Log("memory:",mem.UsedPercent)
		t.Log("memory:",mem.Free)
	})
}

func TestGatewayDownlinkCmd(t *testing.T)  {
	Convey("TestGatewayDownlinkCmd",t, func() {

	})
}