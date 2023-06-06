package services

import (
	"github.com/acework2u/air-iot-app-api-service/repository"
)

type customerService struct {
	cusRepo repository.CustomerRepository
}

func NewCustomerService(cusRepo repository.CustomerRepository) customerService {
	return customerService{cusRepo}
}

func (cs *customerService) CreateNewCustomer(customer *CreateCustomerRequest) (*DBCustomer, error) {

	//var usersResponse *DBCustomer

	cusReq := customer

	res, err := cs.cusRepo.CreateCustomer((*repository.CreateCustomerRequest)(cusReq))
	if err != nil {

		return nil, err
	}

	usersResponse := (*DBCustomer)(res)

	// //log.Println(customer)
	// fmt.Println("In Service")
	// fmt.Println(cusReq)

	//usersResponse.CreateAt = time.Now()

	return usersResponse, nil

}
func (cs *customerService) AllCustomers() ([]*DBCustomer, error) {

	cusRes, err := cs.cusRepo.FindCustomers()

	if err != nil {
		return nil, err
	}

	var customers []*DBCustomer

	for _, customer := range cusRes {
		custRes := &DBCustomer{
			Id:       customer.Id,
			Name:     customer.Name,
			Lastname: customer.Lastname,
			Tel:      customer.Tel,
			Email:    customer.Email,
		}

		customers = append(customers, custRes)

	}

	return customers, nil
}

func (cs *customerService) UpdateCustomer(id string, data *UpdateCustomer) (*DBCustomer, error) {

	cusId := id

	cusData := data

	res, err := cs.cusRepo.UpdateCustomer(cusId, (*repository.UpdateCustomer)(cusData))
	if err != nil {
		return nil, err
	}

	return (*DBCustomer)(res), nil
}
func (cs *customerService) DeleteCustomer(id string) error {

	delID := id

	err := cs.cusRepo.DeleteCustomer(delID)

	if err != nil {
		return err
	}

	return nil
}
