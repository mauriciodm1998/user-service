package service

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

func (c *CustomerRepositoryMock) DeleteCustomer(ctx context.Context, customerId string) error {
	args := c.Called()

	return args.Error(0)
}

// user mock

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

type CustomerServiceMock struct {
	mock.Mock
}

func (c *CustomerServiceMock) CreateCustomer(_ context.Context, customer canonical.Customer) (*canonical.Customer, error) {
	args := c.Called(customer)

	return args.Get(0).(*canonical.Customer), args.Error(1)
}

func (c *CustomerServiceMock) MockCreateCustomer(customer canonical.Customer, errorReturned error, times int) {
	c.On("CreateCustomer", mock.MatchedBy(func(cus canonical.Customer) bool {
		return cus.Document == customer.Document
	})).Return(&customer, errorReturned).Times(times)
}

func (c *CustomerServiceMock) GetCustomer(_ context.Context, customer canonical.Customer) ([]canonical.Customer, error) {
	args := c.Called(customer)

	return args.Get(0).([]canonical.Customer), args.Error(1)
}

func (c *CustomerServiceMock) DeleteCustomer(ctx context.Context, requesterId, customerId string) error {
	args := c.Called()

	return args.Error(0)
}

func (c *CustomerServiceMock) MockGetCustomer(customer canonical.Customer, customers []canonical.Customer, errorReturned error, times int) {
	if customer.Document == "" {
		c.On("GetCustomer", mock.Anything).Return(customers, errorReturned).Times(times)
	} else {
		c.On("GetCustomer", mock.MatchedBy(func(cus canonical.Customer) bool {
			return cus.Document == customer.Document
		})).Return(customers, errorReturned).Times(times)
	}
}

type UserServiceMock struct {
	mock.Mock
}

func (c *UserServiceMock) CreateUser(_ context.Context, customer canonical.Customer, user canonical.User) (*canonical.User, error) {
	args := c.Called(customer, user)

	return args.Get(0).(*canonical.User), args.Error(1)
}

func (c *UserServiceMock) MockCreateUser(user canonical.User, errorReturned error, times int) {
	c.On("CreateUser", mock.Anything, mock.MatchedBy(func(us canonical.User) bool {
		return us.Login == user.Login
	})).Return(&user, errorReturned).Times(times)
}

func (c *UserServiceMock) GetUser(_ context.Context, user canonical.User) ([]canonical.User, error) {
	args := c.Called(user)

	return args.Get(0).([]canonical.User), args.Error(1)
}

func (c *UserServiceMock) MockGetUser(user canonical.User, users []canonical.User, errorReturned error, times int) {
	if user.Login == "" {
		c.On("GetUser", mock.Anything).Return(users, errorReturned).Times(times)
	} else {
		c.On("GetUser", mock.MatchedBy(func(us canonical.User) bool {
			return us.Login == user.Login
		})).Return(users, errorReturned).Times(times)
	}
}
