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
	return nil
}

func (ur *userRepository) Delete(userID uint) error {
	return nil
}

func (ur *userRepository) FindByUsername(username string) (*User, error) {
	return nil, nil
}

func (ur *userRepository) FindByEmail(email string) (*User, error) {
	return nil, nil
}

func (ur *userRepository) FindByID(userID uint) (*User, error) {
	return nil, nil
}

func (ur *userRepository) Update(user *User) error {
	return nil
}
