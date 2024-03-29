package utils

import (
	"encoding/hex"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math"
)

func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}

type powerFunc func(val int) string
type modeFunc func(val int) string
type tempFunc func(val int) string
type roomTempFunc func(val int) string
type setRhFunc func(val int) string
type roomRhFunc func(val int) string
type fanSpeedFunc func(val int) string
type louverFunc func(val int) string
type apsFunc func(val int) string
type ozoneGenFunc func(val int) string

// type pm25Func func(val []byte) float64
type pm25Func func(val []byte) Pm25Info

type AC1000 struct {
	Power    powerFunc    `json:"power"`
	Mode     modeFunc     `json:"mode"`
	Temp     tempFunc     `json:"temp"`
	RoomTemp roomTempFunc `json:"roomTemp"`
	SetRh    setRhFunc    `json:"setRh"`
	RoomRh   roomRhFunc   `json:"roomRh"`
	FanSpeed fanSpeedFunc `json:"fanSpeed"`
	Louver   louverFunc   `json:"louver"`
	Aps      apsFunc      `json:"aps"`
	OzoneGen ozoneGenFunc `json:"ozoneGen"`
	Pm25Info pm25Func     `json:"pm25Info"`
}

type IndoorInfo struct {
	Power    string   `json:"power"`
	Mode     string   `json:"mode"`
	Temp     string   `json:"temp"`
	RoomTemp string   `json:"roomTemp"`
	RhSet    string   `json:"rhSet"`
	RhRoom   string   `json:"RhRoom"`
	FanSpeed string   `json:"fanSpeed"`
	Louver   string   `json:"louver"`
	Aps      string   `json:"aps"`
	OzoneGen string   `json:"ozoneGen"`
	Pm25Info Pm25Info `json:"pm25Info"`
}

type Pm25Info struct {
	Pm25      string `json:"pm25"`
	Clean     string `json:"clean"`
	ResetTime string `json:"resetTime"`
}

type AcValue interface {
	Ac1000() *IndoorInfo
}

type AcStr struct {
	reg1000 []byte
	reg2000 []byte
	reg3000 []byte
	reg4000 []byte
}
type AcValReq struct {
	Reg1000 string
	Reg2000 string
	Reg3000 string
	Reg4000 string
}

func NewGetAcVal(reg *AcValReq) AcValue {
	data, err := hex.DecodeString(reg.Reg1000)
	data2000, err := hex.DecodeString(reg.Reg2000)
	data3000, err := hex.DecodeString(reg.Reg3000)
	data4000, err := hex.DecodeString(reg.Reg4000)
	if err != nil {
		panic(err)
	}

	fmt.Println("ac Val")
	fmt.Println("reg1000 =", data)
	fmt.Println("reg2000 =", data2000)
	fmt.Println("reg3000 =", data3000)
	fmt.Println("reg4000 =", data4000)
	pm25 := data2000[14]
	fmt.Println("pm25", pm25)
	fmt.Println("lne reg1000", len(data))
	fmt.Println("lne reg2000", len(data2000))

	ozoneVal := make([]byte, 2)
	ozoneVal = data2000[18:]
	fmt.Println("data reg2009", data2000[18:])
	fmt.Println("data reg2009", ozoneVal)
	fmt.Println("Prefilter = ", len(data4000))

	fmt.Println("data Pm2.5 low", data2000[14:16])
	fmt.Println("data pm2.5 Hi", data2000[16:18])
	fmt.Println("data pm2.5", data2000[14:18])
	fmt.Println("data pm2.5 = ", GetPm25Val(data2000[14:18]))
	fmt.Printf("data pm2.5 = %.2f \n", GetPm25Val(data2000[14:18]))

	return &AcStr{reg1000: data, reg2000: data2000, reg3000: data3000, reg4000: data4000}
}
func (ut *AcStr) Ac1000() *IndoorInfo {
	ac := &AC1000{
		Power:    power,
		Mode:     mode,
		Temp:     temp,
		RoomTemp: roomTemp,
		SetRh:    rh,
		RoomRh:   rh,
		FanSpeed: fanSpeed,
		Louver:   louver,
		Aps:      Aps,
		OzoneGen: Ozone,
		Pm25Info: Pm25Func,
	}
	rs := &IndoorInfo{
		Power:    ac.Power(int(ut.reg1000[1])),
		Mode:     ac.Mode(int(ut.reg1000[3])),
		Temp:     ac.Temp(int(ut.reg1000[5])),
		RoomTemp: ac.RoomTemp(int(ut.reg1000[7])),
		RhSet:    ac.SetRh(int(ut.reg1000[9])),
		RhRoom:   ac.RoomRh(int(ut.reg1000[11])),
		FanSpeed: ac.FanSpeed(int(ut.reg1000[13])),
		Louver:   ac.Louver(int(ut.reg1000[15])),
		Aps:      ac.Aps(int(ut.reg1000[17])),
		OzoneGen: ac.OzoneGen(int(ut.reg1000[17])),
		Pm25Info: ut.Pm25(),
	}

	return rs
}
func (ut *AcStr) Pm25() Pm25Info {

	cleanTime := ut.reg4000[9]
	pm25Value := GetPm25Val(ut.reg2000[14:18])

	return Pm25Info{
		Pm25:      fmt.Sprintf("%.2f", pm25Value),
		Clean:     fmt.Sprintf("%v", cleanTime),
		ResetTime: fmt.Sprintf("%v", cleanTime),
	}
}

func power(val int) string {
	powerTxt := ""
	switch val {
	case 0:
		powerTxt = "off"
	case 1:
		powerTxt = "on"
	default:
		powerTxt = "err"
	}
	return powerTxt
}
func mode(val int) string {
	displayTxt := ""
	switch val {
	case 0:
		displayTxt = "cool"
	case 1:
		displayTxt = "dry"
	case 2:
		displayTxt = "auto"
	case 3:
		displayTxt = "heat"
	case 4:
		displayTxt = "fan"
	default:
		displayTxt = "err"
	}
	return displayTxt
}
func temp(val int) string {
	displayTxt := ""

	if val < 0 || val > 60 {
		displayTxt = "err"
		return displayTxt
	}
	val2 := float64(val)
	tempVal := val2 / 2
	s := fmt.Sprintf("%3.1f", tempVal)
	displayTxt = s

	return displayTxt

}
func roomTemp(val int) string {
	displayTxt := ""

	if val < 0 || val > 255 {
		displayTxt = "err"
		return displayTxt
	}

	val2 := float64(val)
	tempVal := val2 / 4
	s := fmt.Sprintf("%3.1f", tempVal)
	displayTxt = s

	return displayTxt

}
func rh(val int) string {
	displayTxt := ""

	if val < 0 || val > 100 {
		displayTxt = "err"
		return displayTxt
	}

	displayTxt = fmt.Sprintf("%v", val)

	return displayTxt

}
func fanSpeed(val int) string {
	displayTxt := ""
	//Value 0 : Fan Auto
	//Value 1 : Fan Low
	//Value 2 : Fan Med
	//Value 3 : Fan High
	//Value 4 : Fan Hi Hi
	//Value 5 : Fan Turbo
	switch val {
	case 0:
		displayTxt = "auto"
	case 1:
		displayTxt = "low"
	case 2:
		displayTxt = "med"
	case 3:
		displayTxt = "high"
	case 4:
		displayTxt = "high+"
	default:
		displayTxt = "turbo"
	}
	return displayTxt
}
func louver(val int) string {
	displayTxt := ""
	//Value 0 :  Auto (Swing)
	//Value 1 :  Level 1
	//Value 2 :  Level 2
	//Value 3 :  Level 3
	//Value 4 :  Level 4
	//Value 5 :  Level 5

	switch val {
	case 0:
		displayTxt = "auto"
	case 1:
		displayTxt = "level 1"
	case 2:
		displayTxt = "level 2"
	case 3:
		displayTxt = "level 3"
	case 4:
		displayTxt = "level 4"
	case 5:
		displayTxt = "level 5"
	default:
		displayTxt = "err"
	}
	return displayTxt
}

func Aps(val int) string {
	var displayTxt string
	//
	//fmt.Println("APS=", val)
	//fmt.Println("bits 0=", val&1)
	//fmt.Println("bits 1=", val&2)
	//fmt.Println(val & 4)
	//fmt.Println(val & 8)
	//fmt.Println(val & 16)
	//fmt.Println(val & 32)
	//fmt.Println(val & 64)
	//fmt.Println(val & 128)

	switch val {
	case 0:
	case 32:
		displayTxt = "off"
	case 1:
	case 3:
	case 33:
	case 35:
		displayTxt = "on"
	default:
		displayTxt = "not support"
	}

	return displayTxt
}

func Ozone(val int) string {

	var displayTxt string
	//fmt.Println("bits 0=", val&1)
	//fmt.Println("bits 1=", val&2)
	//fmt.Println(val & 4)
	//fmt.Println(val & 8)
	//fmt.Println(val & 16)
	//fmt.Println(val & 32)
	//fmt.Println(val & 64)
	//fmt.Println(val & 128)

	switch val {
	case 0:
	case 32:
		displayTxt = "off"
	case 1:
	case 33:
		displayTxt = "on"
	case 3:
	case 35:
		displayTxt = "Ozone Generator Running"

	default:
		displayTxt = "this function not support"
	}
	return displayTxt
}
func GetPm25Val(pmReg []byte) float64 {
	if len(pmReg) == 4 {

		var pmVal float64
		for i, v := range pmReg {
			pmVal = pmVal + (float64(v) * math.Pow(2, float64(len(pmReg)-i)))
		}
		// div 1000
		if pmVal != 0 {
			pmVal = pmVal / 1000
		}
		return pmVal
	}
	return 0.0
}

func Pm25Func(pmReg []byte) Pm25Info {
	return Pm25Info{
		Pm25:      fmt.Sprintf("%.2f", GetPm25Val(pmReg[14:18])),
		Clean:     "",
		ResetTime: "",
	}
}
