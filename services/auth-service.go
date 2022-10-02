package services

import (
	"task-5-vix-btpns-mumaralfajar/dto"
	"task-5-vix-btpns-mumaralfajar/models"
	"task-5-vix-btpns-mumaralfajar/repositories"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateToken(user dto.RegisterDto) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	return res
}

func (service *authService) CreateToken(user dto.RegisterDto) models.User {
	newUser := models.User{}
	res := service.userRepository.Insert(newUser)
	return res
}

func (service *authService) FindByEmail(email string) models.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	return service.userRepository.IsDuplicateEmail(email).RowsAffected > 0
}
