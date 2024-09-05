package monitoring

import (
	"testing"
	"time"
)

func TestGetDBInstance(t *testing.T) {
	db1 := GetDBInstance()
	db2 := GetDBInstance()

	if db1 != db2 {
		t.Errorf("Expected same instance, but got different instances")
	}

	if db1.Name != "о" {
		t.Errorf("Expected DB name to be 'о', got %s", db1.Name)
	}
}

func TestDB_RecieveData(t *testing.T) {
	db := GetDBInstance()
	go func() {
		db.Channel <- "Test Data 1"
		db.Channel <- "Test Data 2"
	}()

	db.RecieveData()
}

func TestSensor_Run(t *testing.T) {
	sensor := NewSensor("Temperature")

	go sensor.Run()

	time.Sleep(1 * time.Second)

	select {
	case data := <-sensor.DB.Channel:
		if data == "" {
			t.Error("Expected data from channel, but got empty string")
		}
	case <-time.After(2 * time.Second):
		t.Error("Expected data from sensor, but got timeout")
	}
	sensor.Stop()
}

func TestSensor_Stop(t *testing.T) {
	sensor := NewSensor("Humidity")
	sensor.Run()
	sensor.Stop()

	if sensor.isRun {
		t.Error("Expected sensor to be stopped, but it's still running")
	}
}
