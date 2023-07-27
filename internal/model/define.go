package model

var (
	Post          = "/sys/%s/%s/thing/event/property/post"        //属性上报；productKey,deviceName
	Post_reply    = "/sys/%s/%s/thing/event/property/post_reply"  //上行响应；productKey,deviceName
	Set           = "/sys/%s/%s/thing/service/property/set"       //下行命令;productKey,deviceName
	Set_reply     = "/sys/%s/%s/thing/service/property/set_reply" //下行响应；productKey,deviceName
	Event         = "/sys/%s/%s/thing/event/%s/post"              //事件；productKey,deviceName，${tsl.event.identifier}
	Event_reply   = "/sys/%s/%s/thing/event/%s/post_reply"        //事件响应；productKey,deviceName，${tsl.event.identifier}
	Service       = "/sys/%s/%s/thing/service/%s"                 //服务；productKey,deviceName，${tsl.service.identifier}
	Service_reply = "/sys/%s/%s/thing/service/%s_reply"           //服务响应；服务；productKey,deviceName，${tsl.service.identifier}
)


