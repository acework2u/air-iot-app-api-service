package services

import (
	"errors"
	"fmt"
	"github.com/acework2u/air-iot-app-api-service/repository"
	"time"
)

type deviceService struct {
	deviceRepo repository.DevicesRepository
}

func NewDeviceService(deviceRepo repository.DevicesRepository) DevicesService {
	return &deviceService{deviceRepo: deviceRepo}
}
func (s *deviceService) NewDevice(device *Device) (*ResponseDevice, error) {

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

	respDevice := &ResponseDevice{
		Id:       deviceInfo.Id,
		Name:     deviceInfo.Name,
		SerialNo: deviceInfo.SerialNo,
		Title:    deviceInfo.Title,
		Model:    deviceInfo.Model,
		Warranty: deviceInfo.Warranty,
	}

	return respDevice, nil
}
func (s *deviceService) ListDevice(request *DeviceRequest) ([]*ResponseDevice, error) {
	fmt.Println(request)
	fmt.Println("Device Service Working...")
	var err error
	if len(request.UserId) > 0 {

		filterUid := &repository.DeviceRequest{
			UserId: request.UserId,
		}
		devices, ok := s.deviceRepo.FindDevices(filterUid)

		if ok != nil {
			return nil, ok
		}

		var deviceResponse []*ResponseDevice

		for _, device := range devices {

			deviceRes := &ResponseDevice{
				Id:       device.Id,
				Name:     device.Name,
				SerialNo: device.SerialNo,
				Title:    device.Title,
				Model:    device.Model,
				Warranty: device.Warranty,
			}

			deviceResponse = append(deviceResponse, deviceRes)

		}

		return deviceResponse, nil

	}

	err = errors.New("Don't have data")
	return nil, err
}
