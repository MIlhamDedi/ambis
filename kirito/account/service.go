package account

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"ambis/kirito/pb"
	"ambis/lib/base"
	"ambis/lib/constants"
)

// UserService implements pb.UserServiceServer
type UserService struct {
	UserRepo UserRepo
}

// UserService CreateUser Implementation
func (s *UserService) Create(ctx context.Context, cmd *pb.User) (*pb.Response, error) {
	response := new(pb.Response)
	response.Status = new(pb.Status)
	response.User = new(pb.User)

	userId := uuid.NewV4()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), 8)
	if err != nil {
		response.Status.Success = false
		response.Status.Message = constants.S_INTERNAL_ERROR
		return response, err
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
		response.Status.Success = false
		response.Status.Message = constants.S_INTERNAL_ERROR
		return response, err
	}

	response.User = cmd
	return response, nil
}

// UserService VerifyUser Implementation
func (s *UserService) Verify(ctx context.Context, cmd *pb.User) (*pb.Response, error) {
	response := new(pb.Response)
	response.Status = new(pb.Status)
	response.User = new(pb.User)

	user, err := s.UserRepo.FindByUsername(cmd.Username)
	if err != nil {
		response.Status.Success = false
		response.Status.Message = constants.S_NO_MATCH_FOUND
		return response, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password))
	if err != nil {
		response.Status.Success = false
		response.Status.Message = constants.S_NO_MATCH_FOUND
		return response, err
	}

	response.Status.Success = true
	response.Status.Message = constants.S_NO_ERROR
	response.User.Username = user.Username
	return response, nil
}

// NewUserService instantitae UserService with passed DB as User repository
func NewService(b *base.Base) (*UserService, error) {
	userRepo := UserRepoDB{DB: *b.DB}
	return &UserService{UserRepo: &userRepo}, nil
}
