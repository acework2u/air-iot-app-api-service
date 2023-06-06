package services

type UserResponse struct {
	Code     string `json:"code" bson:"cus_code"`
	Name     string `json:"name" bson:"name"`
	Lastname string `json:"last_name" bson:"last_name"`
	Email    string `json:"email" bson:"email"`
	Mobile   string `json:"mobile_no" bson:"mobile"`
}

type UserService interface {
	GetUsers() ([]*UserResponse, error)
	CreateUser(*UserResponse) (*UserResponse, error)
}
