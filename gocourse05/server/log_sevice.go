package server

import (
	"os"
	"time"
)

type LogService struct {
	strangeApi string // idk just to be...
}

func NewLogService(strangeApi string) *LogService {
	return &LogService{
		strangeApi: strangeApi,
	}
}

func (a *LogService) SendData(s string) {
	f, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write the whole body at once
	currentTime := time.Now()
	_, err = f.WriteString(currentTime.Format(time.RFC3339) + ": Send data to log api " + a.strangeApi + ", data " + s + "\n")
	if err != nil {
		panic(err)
	}
}
