package monitoring

import (
	"fmt"
	"sync"
)

type DB struct {
	Name    string
	Channel chan string
}

var instance *DB
var once sync.Once

func GetDBInstance() *DB {
	once.Do(func() {
		instance = &DB{Name: "о", Channel: make(chan string)}
	})
	return instance
}

func (s *DB) RecieveData() {
	fmt.Print("Start recieving data\n")
	for {
		data, ok := <-s.Channel
		if !ok {

			fmt.Println("Stop receiving data")
			break
		}
		fmt.Println("Received data:", data)
	}
	fmt.Print("Stop recieving data\n")
}
