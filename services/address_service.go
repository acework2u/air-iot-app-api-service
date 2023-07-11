package services

import (
	"github.com/acework2u/air-iot-app-api-service/repository"
)

type addressService struct {
	addrRepo repository.AddressRepository
}

func NewAddressService(addrRepo repository.AddressRepository) AddressService {
	return &addressService{addrRepo}
}

func (s *addressService) NewAddress(address *CustomerAddress) (*DBAddress, error) {

	var userAddress *repository.CustomerAddress = (*repository.CustomerAddress)(address)

	addressRes, err := s.addrRepo.CreateNewAddress(userAddress)

	if err != nil {
		return nil, err
	}
	var responseData *DBAddress = (*DBAddress)(addressRes)

	return responseData, nil
}
