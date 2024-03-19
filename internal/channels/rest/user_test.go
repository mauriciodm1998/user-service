package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"user-service/internal/canonical"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userSvcMock *UserServiceMock
	userRest    User
)

func init() {
	userSvcMock = new(UserServiceMock)

	userRest = NewUserChannel(userSvcMock, customerSvcMock)
}

func TestRestStart(t *testing.T) {
	customerRestMock := &CustomerRestMock{}
	userRestMock := &UserRestMock{}

	restInstance := New(customerSvcMock, userSvcMock)

	customerRestMock.On("RegisterGroup").Once()
	userRestMock.On("RegisterGroup").Once()

	go func() {
		err := restInstance.Start()
		assert.NoError(t, err)
	}()

	time.Sleep(100 * time.Millisecond)
}

func TestUserRegisterGroup(t *testing.T) {
	router := echo.New()

	userRest.RegisterGroup(router.Group("/user"))
	req := createJsonRequest(http.MethodPost, "/user/create", "")
	rec := httptest.NewRecorder()

	router.NewContext(req, rec)
	router.ServeHTTP(rec, req)

	status := rec.Result().StatusCode
	assert.Equal(t, http.StatusBadRequest, status)
}

func TestCreateUser(t *testing.T) {
	user := canonical.User{
		Id:       "asdsadas21",
		Login:    "fulano",
		Password: "12345",
	}

	userSvcMock.MockCreateUser(user, nil, 1)

	request := CreateUserRequest{
		Document: "446.842.868-60",
		Name:     "mauricio Pires",
		Email:    "mauriciodmpires@gmail.com",
		Password: "12345",
		Login:    "fulano",
	}

	req := createJsonRequest(http.MethodPost, "/create", request)
	rec := httptest.NewRecorder()

	err := userRest.Create(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	userSvcMock.AssertExpectations(t)
}

func TestCreateUserBindError(t *testing.T) {
	req := createJsonRequest(http.MethodPost, "/create", "")

	err := userRest.Create(echo.New().NewContext(req, nil))

	assert.Equal(t, "code=400, message={invalid data}", err.Error())
	userSvcMock.AssertExpectations(t)
}

func TestCreateUserServiceError(t *testing.T) {
	userSvcMock.MockCreateUser(canonical.User{
		Login: "123.421.123-60",
	}, errors.New("generic error"), 1)

	request := CreateUserRequest{
		Document: "123.421.123-60",
		Login:    "123.421.123-60",
		Name:     "mauricio Pires",
		Email:    "mauriciodmpires@gmail.com",
		Password: "12345",
	}

	req := createJsonRequest(http.MethodPost, "/create", request)
	rec := httptest.NewRecorder()

	err := userRest.Create(echo.New().NewContext(req, rec))

	assert.Equal(t, "code=500, message={generic error}", err.Error())
	userSvcMock.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	timeNow := time.Now()

	user := canonical.User{
		Id:        "askdjaskdjas",
		Login:     "123.421.123-12",
		Password:  "12345",
		CreatedAt: timeNow,
	}

	userSvcMock.MockGetUser(canonical.User{}, []canonical.User{
		user,
		user,
	}, nil, 1)

	req := createJsonRequest(http.MethodPost, "/", nil)

	rec := httptest.NewRecorder()

	err := userRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	userSvcMock.AssertExpectations(t)
}

func TestGetUserBy(t *testing.T) {
	user := canonical.User{
		Id:        "asdhashjudhas",
		Login:     "123.421.123-12",
		Password:  "12345",
		CreatedAt: time.Now(),
	}

	userSvcMock.MockGetUser(canonical.User{
		Login: "123.421.123-12",
	}, []canonical.User{
		user,
	}, nil, 1)

	req := createJsonRequest(http.MethodPost, "/?login=123.421.123-12", nil)

	rec := httptest.NewRecorder()

	err := userRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	userSvcMock.AssertExpectations(t)
}

func TestGetUserError(t *testing.T) {
	userSvcMock.MockGetUser(canonical.User{}, []canonical.User{}, errors.New("generic error"), 1)

	req := createJsonRequest(http.MethodPost, "/", nil)

	rec := httptest.NewRecorder()

	err := userRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, "code=500, message={generic error}", err.Error())
	userSvcMock.AssertExpectations(t)
}
