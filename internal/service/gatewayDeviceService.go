package service

import (
	"context"
	"deviceservice/internal/conf"
	"deviceservice/internal/model"
	"deviceservice/internal/model/messageBus"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	log "github.com/sirupsen/logrus"
	"runtime"
	"time"
)


/**网关默认的deviceservice服务
 * 默认网关的CPU、内存数据的获取&上传
*/

var (
	GatewayDeviceServiceName = "GatewayDeviceService"
)

type GatewayDeviceService struct {
	Name             string
	ProductKey       string
	DeviceName       string
	Addr             string
	messageBusClient messageBus.MessageBusClient
	recvChan         chan model.ThingsModel
	sendChan         chan model.ThingsModel
}

func GatewayDeviceServerRegister(gateway conf.Gateway) {
	messagebusClient := messageBus.NewMessageBusClient("nanomq")

	mServer := &GatewayDeviceService{Name: GatewayDeviceServiceName,
		messageBusClient: messagebusClient,
		ProductKey: gateway.ProductKey,
		DeviceName: gateway.DeviceName,
		recvChan: make(chan model.ThingsModel, 1024),
		sendChan: make(chan model.ThingsModel, 1024),
	}
	ss = append(ss, mServer)
}

func (gateway *GatewayDeviceService) Start() {
	ctx := context.Background()
	topic := fmt.Sprintf(model.Service, gateway.ProductKey, gateway.DeviceName, "#")
	gateway.messageBusClient.Receive(ctx, topic, gateway.recvChan)
	//todo 接收网关下行命令
	go func() {
		//todo 从messagebus中读取消息，发送到nanomq中
		for {
			timer := time.NewTimer(30 * time.Second)
			select {
			case tm := <-gateway.recvChan:
				log.Debug("gateway recvChan:",tm)
				dispatcher(tm)
				timer.Stop()
			case <-timer.C:
				log.Info("gateway messagebus receive timeout after 5 mins")
				timer.Stop()
			}

		}
	}()
	//todo 定期上报网关数据
	//go func() {
	//
	//}()
}

func (gateway *GatewayDeviceService) Stop() {
	//todo mqtt device server stop
}
func (gateway *GatewayDeviceService) GetName() string {
	return gateway.Name
}

/**
	获取CPU数据
 */
func getCPU() []float64 {
	cpus, _ := cpu.Percent(0, false)
	return cpus
}

/**
	获取内存数据
 */
func getMemory() *mem.VirtualMemoryStat {
	memory, _ := mem.VirtualMemory()
	return memory
}

func getDisk() *disk.UsageStat {
	path := "/"
	if runtime.GOOS == "windows" {
		path = "C:"
	}
	disk, _ := disk.Usage(path)
	return disk
}

func getLoad() *load.AvgStat {
	load, _ := load.Avg()
	return load
}

