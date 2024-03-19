package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"user-service/internal/canonical"

	"github.com/stretchr/testify/assert"
)

var (
	customerRepositoryMock *CustomerRepositoryMock
	customerSvc            CustomerService
)

func init() {
	customerRepositoryMock = new(CustomerRepositoryMock)
	customerSvc = NewCustomerService(customerRepositoryMock)
}

func TestCreateCustomer(t *testing.T) {
	testCases := map[string]struct {
		customer        canonical.Customer
		repositoryError error
		expectedError   error
	}{
		"Sucess": {
			customer: canonical.Customer{
				Document: "44684212301",
				Name:     "mauricio",
				Email:    "tests@email.com",
			},
		},
		"When an error occurred in database": {
			customer: canonical.Customer{
				Document: "44684212300",
				Name:     "mauricio",
				Email:    "tests123@email.com",
			},
			repositoryError: errors.New("generic error"),
			expectedError:   fmt.Errorf("error saving customer in database: %w", errors.New("generic error")),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			customerRepositoryMock.MockCreateCustomer(tc.customer, tc.repositoryError, 1)

			_, err := customerSvc.CreateCustomer(context.Background(), tc.customer)

			assert.Equal(t, tc.expectedError, err)
			customerRepositoryMock.AssertExpectations(t)
		})
	}
}

func TestCustomerGetWithDocument(t *testing.T) {
	customer := canonical.Customer{
		Document: "446",
	}

	customerRepositoryMock.MockGetCustomerByDocument(customer.Document, canonical.Customer{
		Id:       "asdfsa",
		Document: "446",
		Name:     "fulano",
		Email:    "fulano@email",
		UserID:   "sadkjaskdj",
	}, nil, 1)

	customerResponse, err := customerSvc.GetCustomer(context.Background(), customer)

	customerRepositoryMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, customerResponse)
}

func TestGetCustomerWithDocumentError(t *testing.T) {
	customer := canonical.Customer{
		Document: "446",
	}
	customerRepositoryMock.MockGetCustomerByDocument(customer.Document, canonical.Customer{}, errors.New("generic error"), 1)

	_, err := customerSvc.GetCustomer(context.Background(), customer)

	customerRepositoryMock.AssertExpectations(t)
	assert.Equal(t, fmt.Errorf("an error occurred while getting customer in the database: %w", errors.New("generic error")), err)
}

func TestGetCustomerWithUserId(t *testing.T) {
	customer := canonical.Customer{
		UserID: "asdasdas",
	}

	customerRepositoryMock.MockGetCustomerByUserId(customer.UserID, canonical.Customer{
		Id:       "asdfkljasd",
		UserID:   "asdasdas",
		Document: "446",
		Name:     "fulano",
		Email:    "fulano@email",
	}, nil, 1)

	customerResponse, err := customerSvc.GetCustomer(context.Background(), customer)

	customerRepositoryMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, customerResponse)
}

func TestGetCustomerWithUserIdError(t *testing.T) {
	customer := canonical.Customer{
		UserID: "fulano@email",
	}
	customerRepositoryMock.MockGetCustomerByUserId(customer.UserID, canonical.Customer{}, errors.New("generic error"), 1)

	_, err := customerSvc.GetCustomer(context.Background(), customer)

	customerRepositoryMock.AssertExpectations(t)
	assert.Equal(t, fmt.Errorf("an error occurred while getting customer in the database: %w", errors.New("generic error")), err)
}

func TestGetAllCustomers(t *testing.T) {
	customerRepositoryMock.MockGetAllCustomers([]canonical.Customer{
		{
			Id:       "asdfkljasd",
			UserID:   "asdasdas",
			Document: "446",
			Name:     "fulano",
			Email:    "fulano@email",
		},
	}, nil, 1)

	customerResponse, err := customerSvc.GetCustomer(context.Background(), canonical.Customer{})

	customerRepositoryMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, customerResponse)
}

func TestGetAllCustomersError(t *testing.T) {
	customerRepositoryMock.MockGetAllCustomers([]canonical.Customer{}, errors.New("generic error"), 1)

	_, err := customerSvc.GetCustomer(context.Background(), canonical.Customer{})

	customerRepositoryMock.AssertExpectations(t)
	assert.Equal(t, fmt.Errorf("an error occurred while getting customer in the database: %w", errors.New("generic error")), err)
}
