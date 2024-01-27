package mocks

import (
	"context"
	"user-service/internal/canonical"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (c *CustomerRepositoryMock) CreateCustomer(_ context.Context, customer canonical.Customer) error {
	args := c.Called(customer)

	return args.Error(0)
}

func (c *CustomerRepositoryMock) MockCreateCustomer(input canonical.Customer, errRerturn error, times int) {
	c.On("CreateCustomer", mock.MatchedBy(func(cust canonical.Customer) bool {
		return cust.Document == input.Document
	})).Return(errRerturn).Times(times)
}

func (c *CustomerRepositoryMock) GetCustomerByUserId(_ context.Context, email string) (*canonical.Customer, error) {
	args := c.Called(email)

	return args.Get(0).(*canonical.Customer), args.Error(1)
}

func (c *CustomerRepositoryMock) MockGetCustomerByUserId(input string, customerResponse canonical.Customer, errRerturn error, times int) {
	c.On("GetCustomerByUserId", input).Return(&customerResponse, errRerturn).Times(times)
}

func (c *CustomerRepositoryMock) GetCustomerByDocument(_ context.Context, document string) (*canonical.Customer, error) {
	args := c.Called(document)

	return args.Get(0).(*canonical.Customer), args.Error(1)
}

func (c *CustomerRepositoryMock) MockGetCustomerByDocument(input string, customerResponse canonical.Customer, errRerturn error, times int) {
	c.On("GetCustomerByDocument", input).Return(&customerResponse, errRerturn).Times(times)
}

func (c *CustomerRepositoryMock) GetAllCustomers(_ context.Context) ([]canonical.Customer, error) {
	args := c.Called()

	return args.Get(0).([]canonical.Customer), args.Error(1)
}

func (c *CustomerRepositoryMock) MockGetAllCustomers(customerListReturned []canonical.Customer, errorReturned error, times int) {
	c.On("GetAllCustomers").Return(customerListReturned, errorReturned).Times(times)
}

// UserRepositoryMock

// CreateUser(context.Context, canonical.User) error
// GetUserById(context.Context, string) (*canonical.User, error)
// GetUserByLogin(context.Context, string) (*canonical.User, error)
// GetAllUsers(ctx context.Context) ([]canonical.User, error)

type UserRepositoryMock struct {
	mock.Mock
}

func (c *UserRepositoryMock) CreateUser(_ context.Context, User canonical.User) error {
	args := c.Called(User)

	return args.Error(0)
}

func (c *UserRepositoryMock) MockCreateUser(input canonical.User, errRerturn error, times int) {
	c.On("CreateUser", mock.MatchedBy(func(cust canonical.User) bool {
		return cust.Login == input.Login
	})).Return(errRerturn).Times(times)
}

func (c *UserRepositoryMock) GetUserByLogin(_ context.Context, email string) (*canonical.User, error) {
	args := c.Called(email)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (c *UserRepositoryMock) MockGetUserByLogin(input string, UserResponse canonical.User, errRerturn error, times int) {
	c.On("GetUserByLogin", input).Return(&UserResponse, errRerturn).Times(times)
}

func (c *UserRepositoryMock) GetUserById(_ context.Context, id string) (*canonical.User, error) {
	args := c.Called(id)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (c *UserRepositoryMock) MockGetUserById(input string, UserResponse canonical.User, errRerturn error, times int) {
	c.On("GetUserById", input).Return(&UserResponse, errRerturn).Times(times)
}

func (c *UserRepositoryMock) GetAllUsers(_ context.Context) ([]canonical.User, error) {
	args := c.Called()

	return args.Get(0).([]canonical.User), args.Error(1)
}

func (c *UserRepositoryMock) MockGetAllUsers(UserListReturned []canonical.User, errorReturned error, times int) {
	c.On("GetAllUsers").Return(UserListReturned, errorReturned).Times(times)
}
