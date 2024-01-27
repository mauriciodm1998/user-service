package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
	"user-service/internal/canonical"
	"user-service/internal/mocks"
	"user-service/internal/user_grpc_files/user_grpc"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSizeUser = 1024 * 1024

var lisUser *bufconn.Listener

var (
	userServiceMock *mocks.UserServiceMock
)

func init() {
	userServiceMock = new(mocks.UserServiceMock)

	lisUser = bufconn.Listen(bufSizeUser)
	s := grpc.NewServer()
	user_grpc.RegisterUserServiceServer(s, &userGRPCServer{
		UserService: userServiceMock,
	})

	go func() {
		if err := s.Serve(lisUser); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialerUser(context.Context, string) (net.Conn, error) {
	return lisUser.Dial()
}

func TestGetUser(t *testing.T) {
	user := canonical.User{
		Id:        "aksdhaskjd",
		Login:     "11111111",
		Password:  "123",
		CreatedAt: time.Now(),
	}

	user2 := canonical.User{
		Id:        "asmkdhasjhd",
		Login:     "4412",
		Password:  "123",
		CreatedAt: time.Now(),
	}

	userServiceMock.MockGetUser(user, []canonical.User{user, user2}, nil, 1)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialerUser), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := user_grpc.NewUserServiceClient(conn)
	resp, err := client.GetUser(ctx, &user_grpc.User{
		Login: "11111111",
	})

	assert.Equal(t, user.Id, resp.Users[0].Id)
	assert.Equal(t, user.Login, resp.Users[0].Login)
	assert.Equal(t, user.Password, resp.Users[0].Password)

	assert.Equal(t, user2.Id, resp.Users[1].Id)
	assert.Equal(t, user2.Login, resp.Users[1].Login)
	assert.Equal(t, user2.Password, resp.Users[1].Password)

	userServiceMock.AssertExpectations(t)
}

func TestListenUser(t *testing.T) {
	go func() {
		err := ListenUser(8080, userServiceMock)
		if err != nil {
			t.Errorf("Erro ao executar ListenUser: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	serverAddr := "localhost:8080"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer conn.Close()

	assert.Nil(t, err)
	time.Sleep(100 * time.Millisecond)
}
