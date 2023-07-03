package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	Name     string `json:"name"`
	Lastname string `json:"last_name"`
}

type User struct {
	Code     string    `json:"code" bson:"cus_code"`
	Name     string    `json:"name" bson:"name"`
	Lastname string    `json:"last_name" bson:"last_name"`
	Email    string    `json:"email" bson:"email"`
	Mobile   string    `json:"mobile_no" bson:"mobile"`
	CreateAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
type UserDB struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Code     string             `json:"code" bson:"cus_code"`
	Name     string             `json:"name" bson:"name"`
	Lastname string             `json:"last_name" bson:"last_name"`
	Email    string             `json:"email" bson:"email"`
	Mobile   string             `json:"mobile_no" bson:"mobile"`
	CreateAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type UserDB2 struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserSub       string             `json:"usersub" bson:"usersub"`
	Name          string             `json:"name" bson:"name"`
	Lastname      string             `json:"last_name" bson:"last_name"`
	Email         string             `json:"email" bson:"email"`
	Mobile        string             `json:"mobile_no" bson:"mobile"`
	UserConfirmed bool               `json:"UserConfirmed" bson:"UserConfirmed"`
	CreateAt      time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt      time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type CusAddress struct {
	Address  string `json:"address" bson:"address"`
	Statge   string `json:"statge" bson:"statge"`
	District string `json:"district" bson:"district"`
	Province string `json:"provice" bson:"provice"`
	Zipcode  int    `json:"zipcode" bson:"zipcode"`
}

type UserRepository interface {
	CreateUser(*User) (*UserDB, error)
	FindPosts() ([]*UserDB, error)
}

// type PostService interface {
// 	CreatePost(*User) (*User, error)
// 	UpdatePost(string, *User) (*User, error)
// 	FindPostById(string) (*User, error)
// 	FindPosts(page int, limit int) ([]*User, error)
// 	DeletePost(string) error
// }
