package services

import (
	"github.com/mashingan/smapping"
	"log"
	"task-5-vix-btpns-mumaralfajar/dto"
	"task-5-vix-btpns-mumaralfajar/models"
	"task-5-vix-btpns-mumaralfajar/repositories"
)

type UserService interface {
	Update(user dto.UserDto) models.User
	Profile(userID string) models.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) Update(user dto.UserDto) models.User {
	userToUpdate := models.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.Update(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) models.User {
	return service.userRepository.Profile(userID)
}
