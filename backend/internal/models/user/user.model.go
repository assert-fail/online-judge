package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password,omitempty"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Role     string `gorm:"default:user" json:"role"`
}

func (u *User) TableName() string {
	return "user_t"
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}
