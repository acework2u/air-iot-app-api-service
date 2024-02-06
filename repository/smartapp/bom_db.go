package smartapp

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type bomRepositoryDB struct {
	bomCollection *mongo.Collection
	ctx           context.Context
}

func NewBomRepository(ctx context.Context, bomCollection *mongo.Collection) BomRepository {
	return &bomRepositoryDB{ctx: ctx, bomCollection: bomCollection}
}

func (r *bomRepositoryDB) Compressor(indoor string) ([]*AcProduct, error) {

	indNo := strings.ToUpper(indoor)

	filter := bson.M{
		"year":               "$year",
		"odu_model":          "$outdoor_model",
		"odu_item":           "$item_outdoor",
		"ind_model":          "$indoor_model",
		"ind_item":           "$item_indoor",
		"btu":                "$btu",
		"compressor.brand":   "$comp_brand",
		"compressor.model":   "$item_comp_model",
		"compressor.item_no": "$item_comp_code",
	}

	projectStage := bson.D{{"$project", filter}}
	unsetStage := bson.D{{"$unset", bson.A{"_id"}}}
	matchStage := bson.D{{Key: "$match", Value: bson.M{"ind_item": indNo}}}
	pipeline := mongo.Pipeline{projectStage, unsetStage, matchStage}

	cursor, err := r.bomCollection.Aggregate(r.ctx, pipeline)

	//defer cursor.Close(r.ctx)

	if err != nil {
		return nil, err
	}
	var compressors []*AcProduct
	for cursor.Next(r.ctx) {
		compressor := &AcProduct{}
		err := cursor.Decode(compressor)
		if err != nil {
			return nil, err
		}

		compressors = append(compressors, compressor)
	}

	log.Printf("%#v", compressors)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return compressors, nil

}

func (r *bomRepositoryDB) Compressors() ([]*AcProduct, error) {

	filter := bson.M{
		"year":               "$year",
		"odu_model":          "$outdoor_model",
		"odu_item":           "$item_outdoor",
		"ind_model":          "$indoor_model",
		"ind_item":           "$item_indoor",
		"btu":                "$btu",
		"compressor.brand":   "$comp_brand",
		"compressor.model":   "$item_comp_model",
		"compressor.item_no": "$item_comp_code",
	}

	projectStage := bson.D{{"$project", filter}}
	unsetStage := bson.D{{"$unset", bson.A{"_id"}}}
	matchStage := bson.D{{Key: "$match", Value: bson.M{"ind_item": "FSJ-WN009F-DDTGA1"}}}
	pipeline := mongo.Pipeline{projectStage, unsetStage, matchStage}

	cursor, err := r.bomCollection.Aggregate(r.ctx, pipeline)

	//cursor, err := r.bomCollection.Find(r.ctx, query)

	if err != nil {
		return nil, err
	}
	//defer cursor.Close(r.ctx)

	var compressors []*AcProduct

	var res []map[string]interface{}

	for cursor.Next(r.ctx) {
		item := map[string]interface{}{}
		err := cursor.Decode(&item)
		if err != nil {
			return nil, err
		}
		res = append(res, item)
	}

	fmt.Println(res)

	//for cursor.Next(r.ctx) {
	//	compressor := &AcProduct{}
	//	err := cursor.Decode(compressor)
	//	if err != nil {
	//		return nil, err
	//	}
	//	compressors = append(compressors, compressor)
	//}
	//
	//if err := cursor.Err(); err != nil {
	//	return nil, err
	//}
	return compressors, nil
}
