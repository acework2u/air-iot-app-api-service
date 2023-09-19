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

func (s *addressService) UpdateAddress(filter *Filter, addressInfo *CustomerAddress) (*DBAddress, error) {

	return nil, nil
}

func (s *addressService) AllAddress(userid string) ([]*ResponseAddress, error) {

	var myAddress []*ResponseAddress
	resData, err := s.addrRepo.FindAddress(userid)
	if err != nil {
		return nil, err
	}

	for _, cuAddr := range resData {
		addrInfo := &ResponseAddress{
			Id:             cuAddr.Id,
			Name:           cuAddr.Name,
			LastName:       cuAddr.LastName,
			Tel:            cuAddr.Tel,
			Address:        cuAddr.Address,
			District:       cuAddr.District,
			Amphur:         cuAddr.Amphur,
			Province:       cuAddr.Province,
			Zipcode:        cuAddr.Zipcode,
			Tax:            cuAddr.Tax,
			TaxUsed:        cuAddr.Tax_used,
			TaxDefault:     cuAddr.Tax_default,
			AddressDefault: cuAddr.Address_default,
		}

		myAddress = append(myAddress, addrInfo)

	}

	return myAddress, nil
}

func (s *addressService) DelAddress(id string) error {
	//fmt.Printf("Deleted ID %v", id)
	//err := errors.New("Error")
	var err error
	if len(id) > 0 {

		err = s.addrRepo.DeleteAddress(id)
		if err != nil {
			return err
		}
		return nil

	}

	return err
}
