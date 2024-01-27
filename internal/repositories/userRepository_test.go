package repository

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"
	"user-service/internal/canonical"

	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/assert"
)

var (
	userRepo UserRepository
)

func init() {
	userRepo = NewUserRepo(Mock)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	sqlStatement := `INSERT INTO "User" (Id, AccessLevelID, Login, Password, Createdat) VALUES ($1, $2, $3, $4, $5)`

	customer := canonical.User{
		Id:            "askdjask",
		Login:         "fulano@email.com",
		AccessLevelID: 2,
		Password:      "as,asdas",
		CreatedAt:     time.Now(),
	}

	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnResult(pgxmock.NewResult("INSERT", 1))

	err := userRepo.CreateUser(ctx, customer)

	assert.Nil(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestCreateUserError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := `INSERT INTO "User" (Id, AccessLevelID, Login, Password, Createdat) VALUES ($1, $2, $3, $4, $5)`

	customer := canonical.User{
		Id:    "askdjaskdjas",
		Login: "fulano@email.com",
	}

	Mock.ExpectExec(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	err := userRepo.CreateUser(ctx, customer)

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserByLogin(t *testing.T) {
	time := time.Now()
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelID", "CreatedAt"}).
		AddRow("asdasd", "fulano", "221@$123", 2, time))

	customer, err := userRepo.GetUserByLogin(ctx, "fulano@mail")

	assert.Nil(t, err)
	assert.Equal(t, &canonical.User{
		Id:            "asdasd",
		AccessLevelID: 2,
		Login:         "fulano",
		Password:      "221@$123",
		CreatedAt:     time,
	}, customer)

	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserByLoginError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := userRepo.GetUserByLogin(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserByLoginErrorNotFound(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE LOGIN = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelID", "CreatedAt"}))

	_, err := userRepo.GetUserByLogin(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Equal(t, errors.New("entity not found"), err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserById(t *testing.T) {
	time := time.Now()
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE ID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelID", "CreatedAt"}).
		AddRow("asdasd", "123", "221@$123", 2, time))

	customer, err := userRepo.GetUserById(ctx, "asdasd")

	assert.Nil(t, err)
	assert.Equal(t, &canonical.User{
		Id:            "asdasd",
		Login:         "123",
		Password:      "221@$123",
		AccessLevelID: 2,
		CreatedAt:     time,
	}, customer)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserByIdError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE ID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := userRepo.GetUserById(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetUserByIdErrorNotFound(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\" WHERE ID = $1"

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "AccessLevelID", "Login", "Password", "CreatedAt"}))

	_, err := userRepo.GetUserById(ctx, "fulano@mail")

	assert.Error(t, err)
	assert.Equal(t, errors.New("entity not found"), err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetAllUsers(t *testing.T) {
	time := time.Now()
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\""

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnRows(Mock.NewRows([]string{"Id", "Login", "Password", "AccessLevelID", "CreatedAt"}).
		AddRow("asdasd", "logg", "221@$123", 2, time).
		AddRow("asdas", "ligg", "11@$123", 1, time))

	customers, err := userRepo.GetAllUsers(ctx)

	assert.Equal(t, &canonical.User{
		Id:            "asdasd",
		AccessLevelID: 2,
		Login:         "logg",
		Password:      "221@$123",
		CreatedAt:     time,
	}, &customers[0])
	assert.Equal(t, &canonical.User{
		Id:            "asdas",
		Login:         "ligg",
		Password:      "11@$123",
		CreatedAt:     time,
		AccessLevelID: 1,
	}, &customers[1])
	assert.Nil(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}

func TestGetAllUsersError(t *testing.T) {
	ctx := context.Background()
	sqlStatement := "SELECT * FROM \"User\""

	Mock.ExpectQuery(regexp.QuoteMeta(sqlStatement)).WillReturnError(errors.New("generic error"))

	_, err := userRepo.GetAllUsers(ctx)

	assert.Error(t, err)
	assert.Nil(t, Mock.ExpectationsWereMet())
}
