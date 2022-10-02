package services

import (
	"task-5-vix-btpns-mumaralfajar/dto"
	"task-5-vix-btpns-mumaralfajar/models"
	"task-5-vix-btpns-mumaralfajar/repositories"
)

type PhotoService interface {
	Insert(p dto.CreatePhotoDto) models.Photo
	Update(p dto.UpdatePhotoDto) models.Photo
	Delete(p models.Photo)
	All() []models.Photo
	FindById(id string) models.Photo
}

type photoService struct {
	photoRepository repositories.PhotoRepository
}

func NewPhotoService(photoRepository repositories.PhotoRepository) PhotoService {
	return &photoService{
		photoRepository: photoRepository,
	}
}

func (service *photoService) Insert(p dto.CreatePhotoDto) models.Photo {
	photo := models.Photo{}
	photo.Title = p.Title
	photo.Caption = p.Caption
	photo.PhotoUrl = p.PhotoUrl
	photo.UserID = int64(p.UserID)
	res := service.photoRepository.Insert(&photo)
	return res
}

func (service *photoService) Update(p dto.UpdatePhotoDto) models.Photo {
	photo := models.Photo{}
	photo.Title = p.Title
	photo.Caption = p.Caption
	photo.PhotoUrl = p.PhotoUrl
	photo.UserID = int64(p.UserID)
	res := service.photoRepository.Update(&photo)
	return res
}

func (service *photoService) Delete(p models.Photo) {
	service.photoRepository.Delete(p)
}

func (service *photoService) All() []models.Photo {
	return service.photoRepository.All()
}

func (service *photoService) FindById(id string) models.Photo {
	return service.photoRepository.FindById(id)
}
