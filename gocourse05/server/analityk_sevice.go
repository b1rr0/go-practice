package server

import "fmt"

type AnalitykService struct {
	port int
	url  string
}

func NewAnalitykService(port int, url string) *AnalitykService {
	return &AnalitykService{
		port: port,
		url:  url,
	}
}

func (a *AnalitykService) SendData(s string) {
	//send to somewhere
	fmt.Printf("Send data to port %d, url %s, data %s\n", a.port, a.url, s)

}
