package smartapp

type UsersRepository interface {
	Create(user UsersInfo) error
	CreateAddress(userId string, address2 address) error
}

type UserAuth interface {
	SignIn() error
	SignUp(userInfo *UsersInfo) error
	Verified()
}

type address struct {
	ZipCode  int      `bson:"zip_code"`
	Province string   `bson:"province"`
	Amphoe   string   `bson:"amphoe"`
	District string   `bson:"district"`
	Custome1 string   `bson:"custome_1"`
	Default  bool     `bson:"default"`
	Location Location `bson:"location"`
}

type UsersInfo struct {
	Name      string    `bson:"name"`
	LastName  string    `bson:"lastName"`
	addresses []address `bson:"addresses"`
	CreateAt  string    `bson:"createAt"`
	UpdateAt  string    `bson:"updateAt"`
}

type Location struct {
	Lat float32 `bson:"lat"`
	Lan float32 `bson:"lan"`
}
