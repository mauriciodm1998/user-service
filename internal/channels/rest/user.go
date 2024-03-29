package rest

import (
	"fmt"
	"user-service/internal/canonical"
	"user-service/internal/middlewares"
	"user-service/internal/service"

	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	RegisterGroup(*echo.Group)
	Create(echo.Context) error
	Get(ctx echo.Context) error
	HealthCheck(c echo.Context) error
}

type user struct {
	userService service.UserService
}

func NewUserChannel(userService service.UserService, customerSvc service.CustomerService) User {
	return &user{
		userService: userService,
	}
}

func (u *user) RegisterGroup(g *echo.Group) {
	userGroup := g.Group("/user")
	userGroup.Use(middlewares.Authorization)

	userGroup.GET("/", u.Get)
	g.POST("/create", u.Create)
}

func (r *user) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (u *user) Create(c echo.Context) error {
	var userRequest CreateUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, Response{
			Message: fmt.Errorf("invalid data").Error(),
		})
	}

	customer, userTranslated := userRequest.toCanonical()

	user, err := u.userService.CreateUser(c.Request().Context(), customer, userTranslated)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, userToResponse(*user))
}

func (u *user) Get(ctx echo.Context) error {
	queryParams := ctx.QueryParams()

	response, err := u.userService.GetUser(ctx.Request().Context(), canonical.User{
		Login: queryParams.Get("login"),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, Response{
			Message: err.Error(),
		})
	}
	var users []UserResponse

	for _, value := range response {
		users = append(users, userToResponse(value))
	}

	return ctx.JSON(http.StatusOK, users)
}
