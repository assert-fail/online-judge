package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	Delete(userID uint) error
	FindByUsername(username string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByID(userID uint) (*User, error)
	Update(user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) Create(user *User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) Delete(userID uint) error {
	return nil
}

func (ur *userRepository) FindByUsername(username string) (*User, error) {
	var user *User
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindByEmail(email string) (*User, error) {
	var user *User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) FindByID(userID uint) (*User, error) {
	var user *User
	if err := ur.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) Update(user *User) error {
	return nil
}
