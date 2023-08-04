package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
)

const RegisterAddr = int64(1000)
const secretKey = "SaijoDenkiSmartIOT"

type AirCmd interface {
	power() ([]byte, error)
	setTemp() ([]byte, error)
	Action() error
	GetPayload() string
}
type Air struct {
	SerialNo string
	Cmd      string
	Value    string
	Payload  []byte
}

func NewAirCmd(serialNo string, cmd string, value string) AirCmd {
	return &Air{SerialNo: serialNo, Cmd: cmd, Value: value, Payload: []byte{}}
}

func (u *Air) Action() error {

	var err error
	switch strings.ToLower(u.Cmd) {
	case "power":
		u.Payload, err = u.power()
		if err != nil {
			return err
		}

	case "temp":
		u.Payload, err = u.setTemp()
		if err != nil {
			return err
		}

	case "mode":
		u.Payload, err = u.mode()

		if err != nil {
			return err
		}
	case "fan":
		u.Payload, err = u.fan()
		if err != nil {
			return err
		}
	case "swing":
		u.Payload, err = u.swing()
		if err != nil {
			return err
		}
	case "option":
		u.Payload, err = u.option()
		if err != nil {
			return err
		}
	default:

		return errors.New("command is wrong")
	}

	return err
}

func (u *Air) power() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)
	if val > 1 || val < 0 {
		return nil, errors.New("data is wrong")
	}

	regAdd := RegisterAddr
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}

	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)

	var dataFrame []byte = rtuFrame.Bytes()

	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) mode() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)

	if val < 0 || val > 4 {
		return nil, errors.New("value is wrong")
	}
	regAdd := RegisterAddr + 1
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame []byte = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) setTemp() ([]byte, error) {

	tempVal, _ := strconv.ParseFloat(u.Value, 32)
	tempVal = tempVal * 2

	val := uint64(tempVal)

	fmt.Println(val)
	if val < 0 || val > 60 {
		return nil, errors.New("value is wrong")
	}
	regAdd := RegisterAddr + 2
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame []byte = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) fan() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)
	//Value 0 : Fan Auto
	//Value 1 : Fan Low
	//Value 2 : Fan Med
	//Value 3 : Fan High
	//Value 4 : Fan Hi Hi
	//Value 5 : Fan Turbo
	if val < 0 || val > 5 {
		return nil, errors.New("value is wrong")
	}

	regAdd := RegisterAddr + 6
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame []byte = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) swing() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)
	//Value 0 :  Auto (Swing)
	//Value 1 :  Level 1
	//Value 2 :  Level 2
	//Value 3 :  Level 3
	//Value 4 :  Level 4
	//Value 5 :  Level 5
	if val < 0 || val > 5 {
		return nil, errors.New("value is wrong")
	}

	regAdd := RegisterAddr + 7
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame []byte = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) option() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)
	//Value 0 :  Auto (Swing)
	//Value 1 :  Level 1
	//Value 2 :  Level 2
	//Value 3 :  Level 3
	//Value 4 :  Level 4
	//Value 5 :  Level 5
	if val < 0 || val > 5 {
		return nil, errors.New("value is wrong")
	}

	regAdd := RegisterAddr + 8
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame []byte = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) GetPayload() string {

	if len(u.Payload) > 0 {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"serialNumber": string(u.SerialNo),
			"data": map[string]string{
				"cmd": fmt.Sprintf("%x", u.Payload),
			},
		})
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return err.Error()
		}
		return tokenString

	}

	return ""

}

func AirPower(cmd string) *RTUFrame {

	cmdPow := strings.ToLower(cmd)
	cmdVal := 0
	regAdd := int64(RegisterAddr)

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