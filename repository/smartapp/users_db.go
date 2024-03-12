package smartapp

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type userRepositoryDB struct {
	usersCollection *mongo.Collection
	ctx             context.Context
}

func NewUserRepositoryDb(mongo *mongo.Collection, ctx context.Context) UsersRepository {
	return &userRepositoryDB{
		usersCollection: mongo,
		ctx:             ctx,
	}

}

func (r *userRepositoryDB) Create(user UsersInfo) error {
	//panic("no Action")
	currentTime := time.Now()
	user.CreateAt = fmt.Sprintf("", currentTime.Format("2006-01-02 15:04:05"))
	res, err := r.usersCollection.InsertOne(r.ctx, &user)
	_ = res
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 1100 {
			return errors.New("name already exits")
		}
		return err
	}
	opt := options.Index()
	opt.SetUnique(true)

	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
	if _, err := r.usersCollection.Indexes().CreateOne(r.ctx, index); err != nil {
		return errors.New("could not create index for name")
	}

	return nil
}
func (r *userRepositoryDB) CreateAddress(userId string, address2 address) error {
	//panic("no action")
	filter := bson.M{"user": userId}
	_, err := r.usersCollection.Find(r.ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
