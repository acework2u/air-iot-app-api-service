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

		return nil

	case "off":

	}

	return nil
}

func (u *Air) power() ([]byte, error) {
	//cmdPow := strings.ToLower(u.Cmd)
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
