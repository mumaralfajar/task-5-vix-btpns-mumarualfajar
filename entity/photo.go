package entity

type Photo struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"id"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `gorm:"not null" json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
	UserID   int64  `gorm:"not null" json:"user_id"`
	User     User   `gorm:"foreignKey:UserID" json:"user"`
}
