package server

import "gocourse05/server/zoo/photo"

type ZooCameraProcessorServer struct {
	services []GiveNormalNameService
}

func NewZooCameraProcessorServer() *ZooCameraProcessorServer {
	return &ZooCameraProcessorServer{}
}

func (s *ZooCameraProcessorServer) AddService(service GiveNormalNameService) {
	s.services = append(s.services, service)
}

func (s *ZooCameraProcessorServer) ProcessPhoto(photo photo.Photo) {
	for _, service := range s.services {
		service.SendData(photo.PhotoVeiw())
	}
}
