package smartapp

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type errorCodeAcRepo struct {
	errorCodeCollection *mongo.Collection
	ctx                 context.Context
}

func NewErrorCodeRepo(ctx context.Context, errorCodeCollection *mongo.Collection) ErrorCodeRepo {
	return &errorCodeAcRepo{
		ctx:                 ctx,
		errorCodeCollection: errorCodeCollection,
	}
}

func (r *errorCodeAcRepo) ErrorCodeList() ([]*AcErrorCodeDb, error) {
	cursor, err := r.errorCodeCollection.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(r.ctx)

	codeErrors := []*AcErrorCodeDb{}
	for cursor.Next(r.ctx) {
		errCode := &AcErrorCodeDb{}
		err := cursor.Decode(errCode)
		if err != nil {
			return nil, err
		}
		codeErrors = append(codeErrors, errCode)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return codeErrors, nil
}
func (r *errorCodeAcRepo) AcErrCode(code string) (*AcErrorCodeDb, error) {

	filterStage := bson.M{"code": code}
	acErrCode := &AcErrorCodeDb{}
	if err := r.errorCodeCollection.FindOne(r.ctx, filterStage).Decode(acErrCode); err != nil {
		return nil, err
	}
	return acErrCode, nil
}
