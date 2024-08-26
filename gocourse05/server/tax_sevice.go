package server

import "fmt"

type TaxService struct {
	strangeApi string
}

func NewTaxService(strangeApi string) *TaxService {
	return &TaxService{
		strangeApi: strangeApi,
	}
}

func (a *TaxService) SendData(s string) {

	fmt.Printf("Send data to strange api %s, data %s\n", a.strangeApi, s)

}
