package entity

type User struct {
	ID        int64   `gorm:"primary_key:auto_increment" json:"id"`
	Username  string  `gorm:"not null;unique" json:"username"`
	Email     string  `gorm:"not null;unique" json:"email"`
	Password  string  `gorm:"not null" json:"password"`
	Photos    []Photo `gorm:"foreignKey:UserID" json:"photos"`
	CreatedAt string  `gorm:"not null" json:"created_at"`
	UpdatedAt string  `gorm:"not null" json:"updated_at"`
	Token     string  `gorm:"-" json:"token,omitempty"`
}
