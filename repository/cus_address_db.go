package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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
	return nil
}
func (r *AddressRepositoryDB) FindAddress() ([]*DBAddress, error) {
	return nil, nil
}
func (r *AddressRepositoryDB) FindAddressId(string) (*DBAddress, error) {
	return nil, nil
}
