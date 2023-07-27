package service

import (
	"testing"
	"time"
)

func TestMqttStandardDS(t *testing.T)  {
	idMap := make(map[string]interface{})
	idMap["Power"] = struct {
		Value interface{}
		Time time.Time
	}{"on",time.Now()}
	idMap["WF"] = struct {
		Value interface{}
		Time time.Time
	}{23.6,time.Now()}
	//测试数据构造


}
