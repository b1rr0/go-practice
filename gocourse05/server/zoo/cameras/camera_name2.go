package cameras

import (
	"gocourse05/server"
	cameras "gocourse05/server/zoo/photo"
	"time"
)

type CameraNight struct {
	Name     string
	listener server.ZooCameraProcessorServer
}

func NewCameraNight(name string, listener server.ZooCameraProcessorServer) *CameraNight {
	return &CameraNight{
		Name:     name,
		listener: listener,
	}
}

func (c *CameraNight) Run() {
	for {
		c.listener.ProcessPhoto(cameras.RandomNightPhoto("Made on camera:" + c.Name))
		time.Sleep(time.Second)
	}
}
