package services

import "time"

type Device struct {
	Name      string    `json:"name"`
	SerialNo  string    `json:"serialNo"`
	Title     string    `json:"title"`
	Model     string    `json:"model"`
	Warranty  string    `json:"warranty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type DevicesService interface {
	NewDevice() error
	//RegisterDevice()
	//UpdateDevice()
	//DeleteDevice()
	//ConnectDevice()
}
