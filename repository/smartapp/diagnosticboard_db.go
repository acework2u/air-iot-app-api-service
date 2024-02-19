package smartapp

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type diagnosticBoardRepo struct {
	diagnosticBoardCollection *mongo.Collection
	ctx                       context.Context
}

func NewDiagnosticBoardRepo(ctx context.Context, diagnosticCollection *mongo.Collection) DiagnosticBoardRepo {
	return &diagnosticBoardRepo{
		ctx:                       ctx,
		diagnosticBoardCollection: diagnosticCollection,
	}
}
func (r *diagnosticBoardRepo) DiagnosticBoards() ([]DiagnosticBoard, error) {
	query := bson.M{}
	cursor, err := r.diagnosticBoardCollection.Find(r.ctx, query)
	defer cursor.Close(r.ctx)
	if err != nil {
		return nil, err
	}
	diagnosticBoards := []DiagnosticBoard{}
	for cursor.Next(r.ctx) {
		item := DiagnosticBoard{}
		err := cursor.Decode(&item)
		if err != nil {
			return diagnosticBoards, err

		}
		diagnosticBoards = append(diagnosticBoards, item)
	}
	if len(diagnosticBoards) == 0 {
		return diagnosticBoards, mongo.ErrNoDocuments
	}
	return diagnosticBoards, err
}
func (r *diagnosticBoardRepo) DiagnosticBoard(filter DiagnosticFilter) (*DiagnosticBoard, error) {

	query := bson.M{"btu": filter.Btu, "comp_id": filter.CompId}
	diagnosticBoard := &DiagnosticBoard{}

	err := r.diagnosticBoardCollection.FindOne(r.ctx, query).Decode(&diagnosticBoard)
	if err != nil {
		return nil, err
	}
	return diagnosticBoard, err
}
func (r *diagnosticBoardRepo) CompressorOnBoard(filter CompressorFilter) (*DiagnosticBoard, error) {

	query := bson.M{"btu": filter.Btu, "comp_item": filter.CompItem}
	diagnosticBoard := &DiagnosticBoard{}

	err := r.diagnosticBoardCollection.FindOne(r.ctx, query).Decode(&diagnosticBoard)
	if err != nil {
		return nil, err
	}
	return diagnosticBoard, err
}
