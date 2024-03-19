package rest

import (
	"context"
	"user-service/internal/canonical"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

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

func (c *CustomerServiceMock) MockGetCustomer(customer canonical.Customer, customers []canonical.Customer, errorReturned error, times int) {
	if customer.Document == "" {
		c.On("GetCustomer", mock.Anything).Return(customers, errorReturned).Times(times)
	} else {
		c.On("GetCustomer", mock.MatchedBy(func(cus canonical.Customer) bool {
			return cus.Document == customer.Document
		})).Return(customers, errorReturned).Times(times)
	}
}

func (c *CustomerServiceMock) DeleteCustomer(ctx context.Context, requesterId, customerId string) error {
	args := c.Called()

	return args.Error(0)
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

type CustomerRestMock struct {
	mock.Mock
}

type UserRestMock struct {
	mock.Mock
}

func (c *UserRestMock) RegisterGroup(_ *echo.Group) {
}

func (c *UserRestMock) Get(_ echo.Context) error {
	return nil
}

func (c *UserRestMock) Create(_ echo.Context) error {
	return nil
}
