package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
	"user-service/internal/canonical"

	"github.com/stretchr/testify/assert"
	"github.com/undefinedlabs/go-mpatch"
	"golang.org/x/crypto/bcrypt"
)

var (
	userRepositoryMock  *UserRepositoryMock
	customerServiceMock *CustomerServiceMock
	userSvc             UserService
)

func init() {
	userRepositoryMock = new(UserRepositoryMock)
	customerServiceMock = new(CustomerServiceMock)

	userSvc = NewUserService(userRepositoryMock, customerServiceMock)
}

func TestCreate(t *testing.T) {
	testCases := map[string]struct {
		user            canonical.User
		repositoryError error
		expectedError   error
	}{
		"Success": {
			user: canonical.User{
				Id:            "askdjsakjd",
				Login:         "44684212301",
				Password:      "$2a$10$9v97dJUQh.MUDpS69KIRguCP/KVgjEngI/OJ2NlUwrQJhjgPA.5VK",
				AccessLevelID: 2,
				CreatedAt:     time.Now(),
			},
		},
		"When an error occurred in database": {
			user: canonical.User{
				Id:            "12321",
				Login:         "4444444",
				Password:      "$2a$10$asdsa.MUDpS69KIRguCP/KVgjEngI/OJ2NlUwrQJhjgPA.5VK",
				AccessLevelID: 1,
				CreatedAt:     time.Now(),
			},
			repositoryError: errors.New("generic error"),
			expectedError:   fmt.Errorf("error saving user in database: %w", errors.New("generic error")),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			userRepositoryMock.MockCreateUser(tc.user, tc.repositoryError, 1)
			customerServiceMock.MockCreateCustomer(canonical.Customer{
				Document: "123",
			}, nil, 1)

			_, err := userSvc.CreateUser(context.Background(), canonical.Customer{
				Document: "123",
			}, tc.user)

			assert.Equal(t, tc.expectedError, err)
			userRepositoryMock.AssertExpectations(t)
		})
	}
}

func TestCreateUserHashError(t *testing.T) {
	patch, _ := mpatch.PatchMethod(bcrypt.GenerateFromPassword, func(password []byte, cost int) ([]byte, error) {
		return nil, errors.New("generic error")
	})
	defer patch.Unpatch()

	_, err := userSvc.CreateUser(context.Background(), canonical.Customer{}, canonical.User{
		Password: "hasherror",
	})

	assert.Equal(t, fmt.Errorf("error generating password hash: %w", errors.New("generic error")), err)
}

func TestGetUserWithLogin(t *testing.T) {
	user := canonical.User{
		Login: "123",
	}

	userRepositoryMock.MockGetUserByLogin(user.Login, canonical.User{
		Login: "123",
	}, nil, 1)

	userResponse, err := userSvc.GetUser(context.Background(), user)

	userRepositoryMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, userResponse)
}

func TestGetWithLoginError(t *testing.T) {
	user := canonical.User{
		Login: "446",
	}
	userRepositoryMock.MockGetUserByLogin(user.Login, canonical.User{}, errors.New("generic error"), 1)

	_, err := userSvc.GetUser(context.Background(), user)

	userRepositoryMock.AssertExpectations(t)
	assert.Equal(t, fmt.Errorf("an error occurred while getting user in the database: %w", errors.New("generic error")), err)
}

func TestGetAll(t *testing.T) {
	timeNow := time.Now()

	userRepositoryMock.MockGetAllUsers([]canonical.User{
		{
			Id:            "123412",
			Login:         "446",
			Password:      "123124521",
			AccessLevelID: 1,
			CreatedAt:     timeNow,
		},
	}, nil, 1)

	userResponse, err := userSvc.GetUser(context.Background(), canonical.User{})

	userRepositoryMock.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, userResponse)
}

func TestGetAllError(t *testing.T) {
	userRepositoryMock.MockGetAllUsers([]canonical.User{}, errors.New("generic error"), 1)

	_, err := userSvc.GetUser(context.Background(), canonical.User{})

	userRepositoryMock.AssertExpectations(t)
	assert.Equal(t, fmt.Errorf("an error occurred while getting user in the database: %w", errors.New("generic error")), err)
}
