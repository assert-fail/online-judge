package user

import (
	"backend/internal/errors"
	"backend/internal/models/user/request"
	"backend/internal/pkg/logger"
	"backend/internal/pkg/utils"
)

type UserService interface {
	Register(req *request.RegisterBody) error
	Login(req *request.LoginBody) (string, error)
	DeleteUser(userID uint) error
	UpdateUser(user *User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) Register(req *request.RegisterBody) error {
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	if _, err := us.repo.FindByUsername(req.Username); err == nil {
		return errors.NewConflictError("username already exists")
	}

	if _, err := us.repo.FindByEmail(req.Email); err == nil {
		return errors.NewConflictError("email already been registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return us.repo.Create(&user)
}

func (us *userService) Login(req *request.LoginBody) (string, error) {
	foundUser, err := us.repo.FindByUsername(req.Username)
	if err != nil {
		logger.Debug().
			Err(err).
			Str("username", req.Username).
			Msg("User not found by username")
		return "", errors.NewUnauthorizedError("Invalid username or password")
	}

	if err := utils.VerifyPassword(foundUser.Password, req.Password); err != nil {
		logger.Debug().
			Err(err).
			Str("found password", foundUser.Password).
			Str("password", req.Password).
			Msg("Password dose not match")
		return "", errors.NewUnauthorizedError("Invalid username or password")
	}

	token, err := utils.GenerateToken(foundUser.ID, foundUser.Username, foundUser.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *userService) DeleteUser(userID uint) error {
	return nil
}

func (us *userService) UpdateUser(user *User) error {
	return nil
}
