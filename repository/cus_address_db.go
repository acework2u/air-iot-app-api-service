package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var Client *mongo.Client

func init() {
	fmt.Println("address repository")
	err := godotenv.Load()
	//var mongoUrl *string
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err.Error())
		//return os.Getenv("MONGURI")

	}
	dbUrl := os.Getenv("MONGURI")

	Client, err = mongo.NewClient(options.Client().ApplyURI(dbUrl))

	//client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the Database
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}

type AddressRepositoryDB struct {
	addrCollection *mongo.Collection
	ctx            context.Context
	client         *mongo.Client
}

func NewAddressRepositoryDB(addrCollection *mongo.Collection, ctx context.Context) AddressRepository {

	return &AddressRepositoryDB{addrCollection: addrCollection, ctx: ctx, client: Client}
}

func (r *AddressRepositoryDB) CreateNewAddress(address *CustomerAddress) (*DBAddress, error) {

	address.UpdateAt = time.Now()

	res, err := r.addrCollection.InsertOne(r.ctx, address)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("address already exits")
		}
		return nil, err
	}

	var newAddress *DBAddress

	query := bson.M{"_id": res.InsertedID}
	if err = r.addrCollection.FindOne(r.ctx, query).Decode(&newAddress); err != nil {
		return nil, err
	}

	return newAddress, nil
}
func (r *AddressRepositoryDB) UpdateAddress(id string, address *UpdateCustomer) (*DBAddress, error) {
	return nil, nil
}
func (r *AddressRepositoryDB) DeleteAddress(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	res, err := r.addrCollection.DeleteOne(r.ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no address with that Id exists")
	}

	return nil

}
func (r *AddressRepositoryDB) FindAddress(userId string) ([]*DBAddress, error) {

	filter := bson.M{"customerId": userId}

	cursor, err := r.addrCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var myAddress []*DBAddress
	for cursor.Next(r.ctx) {
		addr := &DBAddress{}
		err := cursor.Decode(addr)
		if err != nil {
			return nil, err
		}
		myAddress = append(myAddress, addr)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if len(myAddress) == 0 {
		return []*DBAddress{}, nil
	}

	return myAddress, nil
}
func (r *AddressRepositoryDB) FindAddressId(string) (*DBAddress, error) {
	return nil, nil
}
