package main

import (
	"gocource06/monitoring"
	"time"
)

func main() {
	var s = monitoring.NewSensor("TMP")
	s.Run()
	var s1 = monitoring.NewSensor("TMP1")
	s1.Run()
	var s2 = monitoring.NewSensor("TMP2")
	s2.Run()
	var s3 = monitoring.NewSensor("TMP3")
	s3.Run()

	var listener = monitoring.GetDBInstance()
	go listener.RecieveData()

	time.Sleep(10 * time.Second)
	s.Stop()
	s1.Stop()
	s2.Stop()
	s3.Stop()
}
