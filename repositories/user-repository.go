package repositories

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"task-5-vix-btpns-mumaralfajar/models"
)

type UserRepository interface {
	Insert(user models.User) models.User
	Update(user models.User) models.User
	Profile(user string) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	VerifyCredential(email string, password string) interface{}
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Insert(user models.User) models.User {
	user.Password = db.HashAndSalt([]byte(user.Password))
	db.connection.Create(&user)
	return user
}

func (db *userConnection) Update(user models.User) models.User {
	if user.Password != "" {
		user.Password = db.HashAndSalt([]byte(user.Password))
	} else {
		var oldUser models.User
		db.connection.First(&oldUser, user.ID)
		user.Password = oldUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) Profile(userID string) models.User {
	var user models.User
	db.connection.Preload("Photos").Preload("Photos.User").Find(&user, user.ID)
	return user
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).First(&user)
	return user
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	return db.connection.Where("email = ?", email).First(&models.User{})
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	err := db.connection.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil
	}
	return user
}

// hash and salt
func (db *userConnection) HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
