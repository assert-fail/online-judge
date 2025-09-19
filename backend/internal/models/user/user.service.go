package user

type UserService interface {
	Register(username, email, password, role string) error
	Login(username, password string) (string, error)
	GetUserByID(userID uint) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	DeleteUser(userID uint) error
	UpdateUser(user *User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (us *userService) Register(username, email, password, role string) error {
	return nil
}
func (us *userService) Login(username, password string) (string, error) {
	return "", nil
}
func (us *userService) GetUserByID(userID uint) (*User, error) {
	return nil, nil
}
func (us *userService) GetUserByUsername(username string) (*User, error) {
	return nil, nil
}
func (us *userService) GetUserByEmail(email string) (*User, error) {
	return nil, nil
}
func (us *userService) DeleteUser(userID uint) error {
	return nil
}
func (us *userService) UpdateUser(user *User) error {
	return nil
}
