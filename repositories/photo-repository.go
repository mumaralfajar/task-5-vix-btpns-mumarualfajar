package repositories

import (
	"gorm.io/gorm"
	"task-5-vix-btpns-mumaralfajar/models"
)

type PhotoRepository interface {
	Insert(photo *models.Photo) models.Photo
	Update(photo *models.Photo) models.Photo
	Delete(photo models.Photo)
	All() []models.Photo
	FindById(id string) models.Photo
}

type photoConnection struct {
	connection *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoConnection{
		connection: db,
	}
}

func (db *photoConnection) Insert(photo *models.Photo) models.Photo {
	db.connection.Create(&photo)
	return *photo
}

func (db *photoConnection) Update(photo *models.Photo) models.Photo {
	db.connection.Save(&photo)
	return *photo
}

func (db *photoConnection) Delete(photo models.Photo) {
	db.connection.Delete(&photo)
}

func (db *photoConnection) All() []models.Photo {
	var photos []models.Photo
	db.connection.Preload("User").Find(&photos)
	return photos
}

func (db *photoConnection) FindById(id string) models.Photo {
	var photo models.Photo
	db.connection.Preload("User").Find(&photo, id)
	return photo
}
