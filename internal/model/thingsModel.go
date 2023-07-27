package model

import "time"

type ThingsModel struct {
	Id string
	Version string
	Ack Sys
	Payload interface{}
	Method string
}

type Sys struct {
	Ack bool
}

type Param map[string]Identifer

type Identifer struct {
	Value string
	Time time.Time
}
