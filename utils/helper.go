package utils

import "go.mongodb.org/mongo-driver/bson"

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
type roomTempFunc func(val int) float64
type setRhFunc func(val int) float64
type roomRhFunc func(val int) float64
type fanSpeedFunc func(val int) string
type louverFunc func(val int) string

type AC1000 struct {
	Power    powerFunc    `json:"power"`
	Mode     modeFunc     `json:"mode"`
	Temp     tempFunc     `json:"temp"`
	RoomTemp roomTempFunc `json:"roomTemp"`
	SetRh    setRhFunc    `json:"setRh"`
	RoomRh   roomRhFunc   `json:"roomRh"`
	FanSpeed fanSpeedFunc `json:"fanSpeed"`
	Louver   louverFunc   `json:"louver"`
}

func init() {
	GetAc1000 := &AC1000{
		Power: power,
	}

	_ = GetAc1000

}

func power(val int) string {
	return "off"
}
