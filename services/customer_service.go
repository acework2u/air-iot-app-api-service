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
func (cs *customerService) CreateNewCustomer2(customer *CreateCustomerRequest2) (*DbCustomerResponse2, error) {

	var userResponse *DbCustomerResponse2

	return userResponse, nil

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
func (cs *customerService) UpdateCustomer(id string, data *UpdateInfoRequest) (*DBCustomer, error) {

	response, err := cs.cusRepo.UpdateCustomer(id, (*repository.UpdateCustomer)(data))
	if err != nil {
		return nil, err
	}

	return (*DBCustomer)(response), nil

}
func (cs *customerService) DeleteCustomer(id string) error {

	delID := id

	err := cs.cusRepo.DeleteCustomer(delID)

	if err != nil {
		return err
	}

	return nil
}
func (cs *customerService) CustomerById(uid string) (*repository.DBCustomer2, error) {

	result, err := cs.cusRepo.FindCustomerID(uid)

	if err != nil {
		return nil, err
	}

	//var userInfo *DbCustomerResponse2

	//userInfo.Id = result.Id
	//fmt.Println(result)

	return result, nil
}
