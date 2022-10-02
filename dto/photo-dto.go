package dto

type CreatePhotoDto struct {
	ID       uint64 `json:"id" form:"id"`
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id" binding:"required"`
}

type UpdatePhotoDto struct {
	Title    string `json:"title" form:"title" binding:"required"`
	Caption  string `json:"caption" form:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" form:"photo_url" binding:"required"`
	UserID   uint64 `json:"user_id" form:"user_id" binding:"required"`
}
