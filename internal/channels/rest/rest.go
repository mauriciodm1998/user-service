package rest

import (
	"user-service/internal/config"
	"user-service/internal/middlewares"
	"user-service/internal/service"

	"github.com/labstack/echo/v4"
)

var (
	cfg = &config.Cfg
)

type rest struct {
	customer Customer
	user     User
}

func New(customer service.CustomerService, user service.UserService) rest {
	return rest{
		customer: NewCustomerChannel(customer),
		user:     NewUserChannel(user, customer),
	}
}

func (r rest) Start() error {
	router := echo.New()

	router.Use(middlewares.Logger)

	mainGroup := router.Group("/api")
	mainGroup.GET("/healthz", r.user.HealthCheck)

	customerGroup := mainGroup.Group("/customer")

	customerGroup.Use(middlewares.Authorization)

	r.customer.RegisterGroup(customerGroup)
	r.user.RegisterGroup(mainGroup)

	return router.Start(":" + cfg.Port)
}
