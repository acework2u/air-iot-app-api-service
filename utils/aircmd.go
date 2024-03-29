package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
)

const RegisterAddr = int64(1000)
const Reg2000Addr = int64(2000)
const Reg3000Addr = int64(3000)
const Reg4000Addr = int64(4000)

var secretKey = "SaijoDenkiSmartIOT"

//var secretKey = viper.GetString("SECRET_KEY")

type AirCmd interface {
	power() ([]byte, error)
	setTemp() ([]byte, error)
	Action() error
	GetPayload() string
}
type Air struct {
	SerialNo string `json:"serialNo"`
	Cmd      string `json:"cmd"`
	Value    string `json:"value"`
	Payload  []byte `json:"payload"`
}

func NewAirCmd(serialNo string, cmd string, value string) AirCmd {
	return &Air{SerialNo: serialNo, Cmd: cmd, Value: value, Payload: []byte{}}
}

func (u *Air) Action() error {

	var err error
	switch strings.ToLower(u.Cmd) {
	case "power":
		if strings.ToLower(u.Value) == "on" {
			u.Value = "1"
		}
		if strings.ToLower(u.Value) == "off" {
			u.Value = "0"
		}
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
		switch strings.ToLower(u.Value) {
		case "cool":
		case "0":
			u.Value = "0"
		case "dry":
		case "1":
			u.Value = "1"
		case "auto":
		case "2":
			u.Value = "2"
		case "heat":
		case "3":
			u.Value = "3"
		case "fan":
		case "4":
			u.Value = "4"
		default:
			u.Value = "0"

		}

		u.Payload, err = u.mode()

		if err != nil {
			return err
		}
	case "fan":
		switch strings.ToLower(u.Value) {
		case "auto":
		case "0":
			u.Value = "0"
		case "low":
		case "1":
			u.Value = "1"
		case "med":
		case "2":
			u.Value = "2"
		case "high":
		case "3":
			u.Value = "3"
		case "hihi":
		case "4":
			u.Value = "4"
		case "turbo":
		case "5":
			u.Value = "5"
		default:
			u.Value = ""

		}
		u.Payload, err = u.fan()
		if err != nil {
			return err
		}
	case "swing":
		switch strings.ToLower(u.Value) {
		case "auto":
		case "0":
			u.Value = "0"
		case "level1":
		case "1":
			u.Value = "1"
		case "level2":
		case "2":
			u.Value = "2"
		case "level3":
		case "3":
			u.Value = "3"
		case "level4":
		case "4":
			u.Value = "4"
		case "level5":
		case "5":
			u.Value = "5"
		default:
			u.Value = ""

		}
		u.Payload, err = u.swing()
		if err != nil {
			return err
		}
	case "aps":
		if strings.ToLower(u.Value) == "on" {
			u.Value = "1"
		} else if strings.ToLower(u.Value) == "off" {
			u.Value = "0"
		} else {
			return err
		}

		u.Payload, err = u.ApsControl()
		if err != nil {
			return err
		}
	case "ozone":
		if strings.ToLower(u.Value) == "on" {
			u.Value = "1"
		} else if strings.ToLower(u.Value) == "off" {
			u.Value = "0"
		} else {
			return err
		}

		// set payload
		u.Payload, err = u.OzoneGenerate()
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

	var dataFrame = rtuFrame.Bytes()

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
	var dataFrame = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) setTemp() ([]byte, error) {

	tempVal, _ := strconv.ParseFloat(u.Value, 32)
	tempVal = tempVal * 2

	val := uint64(tempVal)

	//fmt.Println(val)
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
	var dataFrame = rtuFrame.Bytes()
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
	var dataFrame = rtuFrame.Bytes()
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
	var dataFrame = rtuFrame.Bytes()
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
	var dataFrame = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) ApsControl() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)

	if val < 0 || val > 1 {
		return nil, errors.New("value is wrong")
	}

	regAdd := RegisterAddr + 8
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	// enable buzz
	val = val + 32

	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 8)
	payload[3] = uint8(val & 0xff)

	rtuFrame.SetData(payload)
	var dataFrame = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) OzoneGenerate() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)

	if val < 0 || val > 1 {
		return nil, errors.New("value is wrong")
	}

	regAdd := RegisterAddr + 8
	rtuFrame := &RTUFrame{
		Address:  uint8(1),
		Function: uint8(6),
	}
	val = val + 32

	payload := make([]byte, 4)
	payload[0] = uint8(regAdd >> 8)
	payload[1] = uint8(regAdd & 0xff)
	payload[2] = uint8(val >> 7)
	payload[3] = uint8(val & 0xff)
	rtuFrame.SetData(payload)
	var dataFrame = rtuFrame.Bytes()
	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload, nil
}
func (u *Air) ResetTimePreFilter() ([]byte, error) {

	val, _ := strconv.Atoi(u.Value)

	if val < 0 || val > 1 {
		return nil, errors.New("value is wrong")
	}
	regAdd := Reg4000Addr + 8
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
	var dataFrame = rtuFrame.Bytes()
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
func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {

	//secKey := viper.GetString("SECRET_KEY")
	secKey := secretKey

	var secret = []byte(secKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return secret, nil
	})

	if err != nil {

		fmt.Println("Error :" + err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
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

	var dataFrame = rtuFrame.Bytes()

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

	var dataFrame = rtuFrame.Bytes()

	newPayload, _ := NewSaijoFrame(dataFrame)

	return newPayload
}
