package service

import (
	"context"
	"strings"
	"time"

	"github.com/rs401/letsgorip/auth/models"
	"github.com/rs401/letsgorip/auth/repository"
	"github.com/rs401/letsgorip/pb"
	"github.com/rs401/letsgorip/validation"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	usersRepository repository.UsersRepository
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(usersRepository repository.UsersRepository) pb.AuthServiceServer {
	return &authService{usersRepository: usersRepository}
}

func (as *authService) SignUp(ctx context.Context, req *pb.User) (*pb.User, error) {
	err := validation.IsValidSignUp(req)
	if err != nil {
		return nil, err
	}
	exists, err := as.usersRepository.GetByEmail(req.Email)
	if exists.Name != "" {
		return nil, validation.ErrEmailExists
	}

	if exists.Name == "" {
		user := new(models.User)
		user.Name = strings.TrimSpace(req.Name)
		user.Email = validation.NormalizeEmail(req.Email)
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = hash
		user.Role = 0

		err = as.usersRepository.Save(user)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				if strings.Contains(err.Error(), "name") {
					return nil, validation.ErrNameExists
				}
				if strings.Contains(err.Error(), "email") {
					return nil, validation.ErrEmailExists
				}
			}
			return nil, err
		}
		return user.ToProtoBuffer(), nil
	}

	return nil, err

}

func (as *authService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.User, error) {
	user, err := as.usersRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return user.ToProtoBuffer(), nil
}

func (as *authService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	user, err := as.usersRepository.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return user.ToProtoBuffer(), nil
}

func (as *authService) ListUsers(req *pb.ListUsersRequest, stream pb.AuthService_ListUsersServer) error {
	users, err := as.usersRepository.GetAll()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := stream.Send(user.ToProtoBuffer())
		if err != nil {
			return err
		}
	}
	return nil
}

func (as *authService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	// Verify user exists
	user, err := as.usersRepository.GetById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, validation.ErrNotFound
	}

	// Validate user name not empty
	if validation.IsEmptyString(req.Name) {
		return nil, validation.ErrEmptyName
	}

	// Validate user email not empty
	if validation.IsEmptyString(req.Email) {
		return nil, validation.ErrEmptyEmail
	}

	// Validate user email is email
	if !validation.IsValidEmail(req.Email) {
		return nil, validation.ErrInvalidEmail
	}

	// Update the user record
	user.Name = req.Name
	user.Email = req.Email
	user.UpdatedAt = time.Now()

	err = as.usersRepository.Update(user)
	return user.ToProtoBuffer(), err

}

func (as *authService) DeleteUser(ctx context.Context, req *pb.GetUserRequest) (*pb.DeleteUserResponse, error) {
	err := as.usersRepository.Delete(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResponse{Id: req.Id}, nil
}
