package smartapp

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DiagnosticBoardRepo interface {
	DiagnosticBoards() ([]DiagnosticBoard, error)
	DiagnosticBoard(filter DiagnosticFilter) (*DiagnosticBoard, error)
}

type DiagnosticFilter struct {
	Btu    int64  `json:"btu"`
	CompId string `json:"comp_id"`
}

type DiagnosticBoard struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Btu       int64              `bson:"btu" json:"btu"`
	CompBrand string             `bson:"comp_brand" json:"comp_brand"`
	CompId    string             `bson:"comp_id" json:"comp_id"`
	CompItem  string             `bson:"comp_item" json:"comp_item"`
	CompModel string             `bson:"comp_model" json:"comp_model"`
}
