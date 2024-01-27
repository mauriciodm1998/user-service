package rest

import (
	"user-service/internal/canonical"
	"user-service/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type Customer interface {
	RegisterGroup(*echo.Group)
	Get(ctx echo.Context) error
}

type customer struct {
	service service.CustomerService
}

func NewCustomerChannel(service service.CustomerService) Customer {
	return &customer{
		service: service,
	}
}

func (u *customer) RegisterGroup(g *echo.Group) {
	g.GET("/", u.Get)
}

func (u *customer) Get(ctx echo.Context) error {
	queryParams := ctx.QueryParams()

	response, err := u.service.GetCustomer(ctx.Request().Context(), canonical.Customer{
		UserID:   queryParams.Get("userid"),
		Document: queryParams.Get("document"),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	var users []CustomerResponse

	for _, value := range response {
		users = append(users, customerToResponse(value))
	}

	return ctx.JSON(http.StatusOK, users)
}
