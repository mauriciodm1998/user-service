package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"user-service/internal/canonical"
	"user-service/internal/mocks"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	customerSvcMock *mocks.CustomerServiceMock
	customerRest    Customer
)

func init() {
	customerSvcMock = new(mocks.CustomerServiceMock)

	customerRest = NewCustomerChannel(customerSvcMock)
}

func TestCustomerRegisterGroup(t *testing.T) {
	router := echo.New()

	channel := NewCustomerChannel(customerSvcMock)
	channel.RegisterGroup(router.Group("/customer"))

	req := createJsonRequest(http.MethodGet, "/customer/123", nil)
	rec := httptest.NewRecorder()

	router.NewContext(req, rec)
	router.ServeHTTP(rec, req)

	status := rec.Result().StatusCode
	assert.Equal(t, http.StatusNotFound, status)
}

func TestGetCustomer(t *testing.T) {
	customer := canonical.Customer{
		Id:       "asdasdas",
		Document: "123.421.123-12",
		Name:     "mauricio Pires",
		Email:    "mauriciodmpires@gmail.com",
	}

	customerSvcMock.MockGetCustomer(canonical.Customer{}, []canonical.Customer{
		customer,
		customer,
	}, nil, 1)

	req := createJsonRequest(http.MethodPost, "/", nil)

	rec := httptest.NewRecorder()

	err := customerRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	customerSvcMock.AssertExpectations(t)
}

func TestGetCustomerBy(t *testing.T) {
	customer := canonical.Customer{
		Id:       "asdasdasdfas",
		Document: "123.421.123-12",
		Name:     "mauricio Pires",
		Email:    "mauriciodmpires@gmail.com",
	}

	customerSvcMock.MockGetCustomer(canonical.Customer{
		Document: "123.421.123-12",
	}, []canonical.Customer{
		customer,
	}, nil, 1)

	req := createJsonRequest(http.MethodPost, "/?document=123.421.123-12", nil)

	rec := httptest.NewRecorder()

	err := customerRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	assert.Nil(t, err)
	customerSvcMock.AssertExpectations(t)
}

func TestGetCustomerError(t *testing.T) {
	customerSvcMock.MockGetCustomer(canonical.Customer{}, []canonical.Customer{}, errors.New("generic error"), 1)

	req := createJsonRequest(http.MethodPost, "/", nil)

	rec := httptest.NewRecorder()

	err := customerRest.Get(echo.New().NewContext(req, rec))

	assert.Equal(t, "code=500, message={generic error}", err.Error())
	customerSvcMock.AssertExpectations(t)
}

func createJsonRequest(method, endpoint string, request interface{}) *http.Request {
	json, _ := json.Marshal(request)
	req := httptest.NewRequest(method, endpoint, bytes.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	return req
}
