package services

import (
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
)

type addressService struct {
	addrRepo repository.AddressRepository
}

func NewAddressService(addrRepo repository.AddressRepository) AddressService {
	return &addressService{addrRepo}
}

func (s *addressService) NewAddress(address *CustomerAddress) (*DBAddress, error) {

	fmt.Println(address)

	return nil, nil
}
