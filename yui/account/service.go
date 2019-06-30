package account

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserService interface {
		CreateUser(cmd *CreateUser) (string, error)
		VerifyUser(cmd *VerifyUser) (*User, error)
	}

	UserServiceImpl struct {
		UserRepo UserRepo
	}
)

// UserService CreateUser Implementation
func (s *UserServiceImpl) CreateUser(cmd *CreateUser) (string, error) {
	userId := uuid.NewV4()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 8)
	if err != nil {
		return "", err
	}

	user := User{
		ID:        userId.String(),
		Username:  cmd.Username,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Password:  string(hashedPassword),
	}

	err = s.UserRepo.Save(&user)

	if err != nil {
		return "", err
	}

	return user.Username, nil
}

// UserService VerifyUser Implementation
func (s *UserServiceImpl) VerifyUser(cmd *VerifyUser) (*User, error) {
	user, err := s.UserRepo.FindByUsername(cmd.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
