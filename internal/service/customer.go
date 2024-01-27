package service

import (
	"context"
	"fmt"
	"user-service/internal/canonical"
	repository "user-service/internal/repositories"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CustomerService interface {
	CreateCustomer(context.Context, canonical.Customer) (*canonical.Customer, error)
	GetCustomer(ctx context.Context, customer canonical.Customer) ([]canonical.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	return &customerService{repository}
}

func (u *customerService) CreateCustomer(ctx context.Context, customer canonical.Customer) (*canonical.Customer, error) {
	customer.Id = uuid.New().String()

	err := u.repo.CreateCustomer(ctx, customer)
	if err != nil {
		err = fmt.Errorf("error saving customer in database: %w", err)
		logrus.WithError(err).Warn()
		return nil, err
	}

	return &customer, nil
}

func (u *customerService) GetCustomer(ctx context.Context, customer canonical.Customer) ([]canonical.Customer, error) {
	var response []canonical.Customer

	if customer.Document != "" {
		baseCustomer, err := u.repo.GetCustomerByDocument(ctx, customer.Document)
		if err != nil {
			err = fmt.Errorf("An error occurred while getting customer in the database: %w", err)
			logrus.WithFields(logrus.Fields{"customer Document:": customer.Document}).Error(err)
			return nil, err
		}

		response = append(response, *baseCustomer)

		return response, nil
	}

	if customer.UserID != "" {
		baseCustomer, err := u.repo.GetCustomerByUserId(ctx, customer.UserID)
		if err != nil {
			err = fmt.Errorf("An error occurred while getting customer in the database: %w", err)
			logrus.WithFields(logrus.Fields{"customer email:": customer.Email}).Error(err)
			return nil, err
		}

		response = append(response, *baseCustomer)

		return response, nil
	}

	response, err := u.repo.GetAllCustomers(ctx)
	if err != nil {
		err = fmt.Errorf("An error occurred while getting customer in the database: %w", err)
		logrus.Error(err)
		return nil, err
	}

	return response, nil
}
