package mocks

import (
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type CustomerRestMock struct {
	mock.Mock
}

func (c *CustomerRepositoryMock) RegisterGroup(_ *echo.Group) {
}

func (c *CustomerRepositoryMock) Get(_ echo.Context) error {
	return nil
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
