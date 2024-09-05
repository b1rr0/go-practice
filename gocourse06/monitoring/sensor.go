package monitoring

import (
	"fmt"
	"time"
)

type Sensor struct {
	Type  string
	DB    *DB
	isRun bool
}

func NewSensor(sensorType string) Sensor {
	return Sensor{
		Type:  sensorType,
		DB:    GetDBInstance(),
		isRun: false,
	}
}

func (s *Sensor) Run() {
	s.isRun = true
	go func() {
		for s.isRun {
			select {
			case s.DB.Channel <- fmt.Sprintf("%s: шось кудись пішло шось там помірялося", s.Type):
				fmt.Print(s.Type, " Send to channel\n")
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
}

func (s *Sensor) Stop() {
	s.isRun = false
}

func (s *Sensor) GenerateDate() {
	s.DB.RecieveData()
}
