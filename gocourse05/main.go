package main

import (
	"gocourse05/server"
	"gocourse05/server/zoo/cameras"
)

func main() {

	listener := server.NewZooCameraProcessorServer()
	listener.AddService(server.NewAnalitykService(555, "2323"))
	listener.AddService(server.NewTaxService("BOOO"))

	var camerasArr []cameras.Camera

	camerasArr = append(camerasArr, cameras.NewCameraDay("DayCam1", *listener))
	camerasArr = append(camerasArr, cameras.NewCameraDay("DayCam2", *listener))
	camerasArr = append(camerasArr, cameras.NewCameraDay("DayCam3", *listener))
	camerasArr = append(camerasArr, cameras.NewCameraNight("NightCam1", *listener))
	camerasArr = append(camerasArr, cameras.NewCameraNight("NightCam2", *listener))
	camerasArr = append(camerasArr, cameras.NewCameraNight("NightCam2", *listener))

	for _, camera := range camerasArr {
		camera.Run()
	}

}
