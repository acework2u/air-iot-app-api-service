package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepositoryDB struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserRepositoryDB(mongo *mongo.Collection, ctx context.Context) userRepositoryDB {
	return userRepositoryDB{userCollection: mongo, ctx: ctx}
}

func (r *userRepositoryDB) CreateUser(user *User) (*UserDB, error) {
	return nil, nil
}
func (r *userRepositoryDB) FindPosts() ([]*UserDB, error) {

	query := bson.M{}
	cursor, err := r.userCollection.Find(r.ctx, query)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(r.ctx)

	var response []*UserDB

	for cursor.Next(r.ctx) {
		userInfo := &UserDB{}
		err := cursor.Decode(userInfo)
		if err != nil {
			return nil, err
		}

		response = append(response, userInfo)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return []*UserDB{}, nil
	}

	return response, nil
}

/*

func NewUserRepositoryDB(mongo *mongo.Collection, ctx context.Context) UserRepository {
	return &UserRepositoryDB{userCollection: mongo, ctx: ctx}
}

func (r *UserRepositoryDB) CreateUser(user *User) (*UserDB, error) {

	fmt.Println("Model DB---------->")

	user.CreateAt = time.Now()
	user.UpdateAt = user.CreateAt

	res, err := r.userCollection.InsertOne(r.ctx, user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that name already exits")
		}
	}
	return nil, err

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"name": 1}, Options: opt}
	if _, err := r.userCollection.Indexes().CreateOne(r.ctx, index); err != nil {
		return nil, errors.New("Could not create index fot name")
	}
	var newUser *UserDB
	query := bson.M{"_id": res.InsertedID}
	if err = r.userCollection.FindOne(r.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (r *UserRepositoryDB) FindPosts(user *User) (*User, error) {

	return user, nil
}

*/
