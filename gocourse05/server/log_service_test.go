package server

import (
	"os"
	"testing"
	"time"
)

func TestSendData(t *testing.T) {
	a := NewLogService("test")
	a.SendData("test data")

	expected := time.Now().Format(time.RFC3339) + ": Send data to log api test, data test data\n"
	actual, err := os.ReadFile("log.txt")
	if err != nil {
		t.Error(err)
	}
	if string(actual) != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
