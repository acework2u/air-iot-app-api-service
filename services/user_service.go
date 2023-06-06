package services

import (
	"log"

	"github.com/acework2u/air-iot-app-api-service/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) userService {
	return userService{userRepo: userRepo}
}

func (s *userService) GetUsers() ([]*UserResponse, error) {

	users, err := s.userRepo.FindPosts()

	if err != nil {
		log.Println(err)
		return nil, err

	}
	var usersResponse []*UserResponse

	for _, user := range users {

		userRes := &UserResponse{
			Code:     user.Code,
			Name:     user.Name,
			Lastname: user.Lastname,
			Email:    user.Email,
			Mobile:   user.Mobile,
		}

		usersResponse = append(usersResponse, userRes)

	}

	return usersResponse, nil
}

func (s *userService) CreateUser(customer *UserResponse) (*UserResponse, error) {
	return nil, nil
}
