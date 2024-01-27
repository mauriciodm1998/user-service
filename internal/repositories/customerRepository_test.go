package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"user-service/internal/canonical"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

var (
	customerRepo CustomerRepository
	Mock         pgxmock.PgxPoolIface
)

func init() {
	mock, _ := pgxmock.NewPool()

	customerRepo = NewCustomerRepo(mock)

	Mock = mock
}

func TestCreateCustomer(t *testing.T) {
	ctx := context.Background()
	sqlStatement := `INSERT INTO "Customer" (Id, UserID, Document, Name, Email) VALUES ($1, $2, $3, $4, $5)`

	customer := canonical.Customer{
		Id:    "sajdhasjkdh",
		Email: "fulano@email.com",
	}

	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := customerRepo.CreateCustomer(ctx, customer)

	assert.Nil(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestCreateCustomerError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := `INSERT INTO "Customer" (Id, UserID, Document, Name, Email) VALUES ($1, $2, $3, $4, $5)`

	customer := canonical.Customer{
		Id:    "asdkljaskdjask",
		Email: "fulano@email.com",
	}

	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	err := customerRepo.CreateCustomer(ctx, customer)

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByUserId(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE UserID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "UserId", "Document", "Name", "Email"}).
		AddRow("asdsad", "asdfasdf", "123", "fulanos", "fulano@mail"))

	customer, err := customerRepo.GetCustomerByUserId(ctx, "fulano@mail")

	assert.Nil(t, err)
	assert.Equal(t, &canonical.Customer{
		Id:       "asdsad",
		Document: "123",
		Name:     "fulanos",
		Email:    "fulano@mail",
		UserID:   "asdfasdf",
	}, customer)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByUserIdError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE UserID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := customerRepo.GetCustomerByUserId(ctx, "askdjaskj")

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByUserIdErrorNotFound(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE UserID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "UserId", "Document", "Name", "Email"}))

	_, err := customerRepo.GetCustomerByUserId(ctx, "asdjasdk-asdjkjas")

	assert.Error(t, err)
	assert.Equal(t, errors.New("entity not found"), err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByDocument(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE Document = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "UserId", "Document", "Name", "Email"}).
		AddRow("askdjsakdjas", "askdjsakdjasasd", "123", "fulanos", "fulano@mail"))

	customer, err := customerRepo.GetCustomerByDocument(ctx, "fulano@mail")

	assert.Nil(t, err)
	assert.Equal(t, &canonical.Customer{
		Id:       "askdjsakdjas",
		Document: "123",
		Name:     "fulanos",
		Email:    "fulano@mail",
		UserID:   "askdjsakdjasasd",
	}, customer)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByDocumentError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE Document = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := customerRepo.GetCustomerByDocument(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetCustomerByDocumentErrorNotFound(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\" WHERE Document = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "UserId", "Document", "Name", "Email"}))

	_, err := customerRepo.GetCustomerByDocument(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Equal(t, errors.New("entity not found"), err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetAllCustomers(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\""

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "UserId", "Document", "Name", "Email"}).
		AddRow("askdjsakdjas", "askdjsakdjasasd", "123", "fulanos", "fulano@mail").
		AddRow("asdasd", "asdsva", "431", "fulana", "fulana@mail"))

	customers, err := customerRepo.GetAllCustomers(ctx)

	assert.Equal(t, &canonical.Customer{
		Id:       "askdjsakdjas",
		UserID:   "askdjsakdjasasd",
		Document: "123",
		Name:     "fulanos",
		Email:    "fulano@mail",
	}, &customers[0])
	assert.Equal(t, &canonical.Customer{
		Id:       "asdasd",
		UserID:   "asdsva",
		Document: "431",
		Name:     "fulana",
		Email:    "fulana@mail",
	}, &customers[1])
	assert.Nil(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetAllCustomersError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"Customer\""

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := customerRepo.GetAllCustomers(ctx)

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}
