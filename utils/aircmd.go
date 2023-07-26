package utils

import (
	"strings"
)

func AirPower(cmd string) *RTUFrame {

	cmdPow := strings.ToLower(cmd)
	cmdVal := 0
	regAdd := int64(1000)

	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}

	//buff := new(bytes.Buffer)

	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)

	if cmdPow == "on" {
		cmdVal = 1
		payload[2] = uint8(cmdVal >> 8)
		payload[3] = uint8(cmdVal & 0xff)
	}
	if cmdPow == "off" {
		cmdVal = 0
		payload[2] = uint8(cmdVal >> 8)
		payload[3] = uint8(cmdVal & 0xff)
	}
	rtuFrame.SetData(payload)

	var dataFrame []byte = rtuFrame.Bytes()

	newPayload, _ := NewRTUFrame(dataFrame)

	return newPayload
}
func AirPower2(cmd string) []byte {

	cmdPow := strings.ToLower(cmd)
	cmdVal := 0
	regAdd := int64(1000)

	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}

	//buff := new(bytes.Buffer)

	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)

	if cmdPow == "on" {
		cmdVal = 1
		payload[2] = uint8(cmdVal >> 8)
		payload[3] = uint8(cmdVal & 0xff)
	}
	if cmdPow == "off" {
		cmdVal = 0
		payload[2] = uint8(cmdVal >> 8)
		payload[3] = uint8(cmdVal & 0xff)
	}
	rtuFrame.SetData(payload)

	var dataFrame []byte = rtuFrame.Bytes()

	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload
}
