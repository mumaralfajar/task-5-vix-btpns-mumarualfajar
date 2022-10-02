package entity

type User struct {
	ID        int64   `gorm:"primary_key:auto_increment" json:"id"`
	Username  string  `gorm:"not null;unique" json:"username"`
	Email     string  `gorm:"not null;unique" json:"email"`
	Password  string  `gorm:"->;<-;not null" json:"-"`
	Photo     []Photo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"photo"`
	CreatedAt string  `gorm:"not null" json:"created_at"`
	UpdatedAt string  `gorm:"not null" json:"updated_at"`
	Token     string  `gorm:"-" json:"token,omitempty"`
}
