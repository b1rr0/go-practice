package server

import (
	"os"
	"strings"
	"testing"
	"time"
)

func TestSendData(t *testing.T) {
	a := NewLogService("test")
	a.SendData("test data")

	expected := time.Now().Format(time.RFC3339) + ": Send data to log api test, data test data\n"
	actual, err := os.ReadFile("log.txt")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasSuffix(string(actual), expected) {
		t.Fatalf("expected suffix %q, got %q", expected, actual)
	}
}
