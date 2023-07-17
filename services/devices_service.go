package services

import (
	"github.com/acework2u/air-iot-app-api-service/repository"
	"time"
)

type deviceService struct {
	deviceRepo repository.DevicesRepository
}

func NewDeviceService(deviceRepo repository.DevicesRepository) DevicesService {
	return &deviceService{deviceRepo: deviceRepo}
}

func (s *deviceService) NewDevice(device *Device) (*responseDevice, error) {

	//fmt.Println(device)
	regTime := time.Now()
	deviceDB := &repository.Device{
		Name:      device.Name,
		Title:     device.Title,
		Model:     device.Model,
		UserId:    device.UserId,
		SerialNo:  device.SerialNo,
		Warranty:  device.Warranty,
		CreatedAt: regTime,
		UpdatedAt: regTime,
	}

	deviceInfo, err := s.deviceRepo.CreateDevice(deviceDB)
	if err != nil {
		return nil, err
	}

	respDevice := &responseDevice{
		Id:       deviceInfo.Id,
		Name:     deviceInfo.Name,
		SerialNo: deviceInfo.SerialNo,
		Title:    deviceInfo.Title,
		Model:    deviceInfo.Model,
		Warranty: deviceInfo.Warranty,
	}

	return respDevice, nil
}
