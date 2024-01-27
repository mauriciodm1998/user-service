package service

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/canonical"
	repository "user-service/internal/repositories"
	"user-service/internal/security"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, customer canonical.Customer, user canonical.User) (*canonical.User, error)
	GetUser(ctx context.Context, user canonical.User) ([]canonical.User, error)
}

type userService struct {
	customerService CustomerService
	userRepo        repository.UserRepository
}

func NewUserService(repo repository.UserRepository, customerService CustomerService) UserService {
	return &userService{
		userRepo:        repo,
		customerService: customerService,
	}
}

func (u *userService) CreateUser(ctx context.Context, customer canonical.Customer, user canonical.User) (*canonical.User, error) {
	user.CreatedAt = time.Now()
	passEncrypted, err := security.Hash(user.Password)
	if err != nil {
		err = fmt.Errorf("error generating password hash: %w", err)
		logrus.WithError(err).Warn()
		return nil, err
	}
	user.Password = string(passEncrypted)

	user.Id = uuid.New().String()
	user.AccessLevelID = 1
	err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		err = fmt.Errorf("error saving user in database: %w", err)
		logrus.WithError(err).Warn()
		return nil, err
	}

	customer.UserID = user.Id

	_, err = u.customerService.CreateCustomer(ctx, customer)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userService) GetUser(ctx context.Context, user canonical.User) ([]canonical.User, error) {
	var response []canonical.User

	if user.Login != "" {
		baseUser, err := u.userRepo.GetUserByLogin(ctx, user.Login)
		if err != nil {
			err = fmt.Errorf("An error occurred while getting user in the database: %w", err)
			logrus.WithFields(logrus.Fields{"user Document:": user.Login}).Error(err)
			return nil, err
		}

		response = append(response, *baseUser)

		return response, nil
	}

	response, err := u.userRepo.GetAllUsers(ctx)
	if err != nil {
		err = fmt.Errorf("An error occurred while getting user in the database: %w", err)
		logrus.Error(err)
		return nil, err
	}

	return response, nil
}
