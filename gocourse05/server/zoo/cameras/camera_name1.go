package cameras

import (
	"gocourse05/server"
	cameras "gocourse05/server/zoo/photo"
	"time"
)

type CameraDay struct {
	Name     string
	listener server.ZooCameraProcessorServer
}

func NewCameraDay(name string, listener server.ZooCameraProcessorServer) *CameraDay {
	return &CameraDay{
		Name:     name,
		listener: listener,
	}
}

func (c *CameraDay) Run() {

	for {
		c.listener.ProcessPhoto(cameras.RandomDayPhoto("Made on camera:" + c.Name))
		time.Sleep(time.Second)
	}

}
