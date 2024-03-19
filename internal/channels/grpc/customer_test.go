package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"
	"user-service/internal/canonical"
	"user-service/internal/customer_grpc_files/customer_grpc"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var (
	customerServiceMock *CustomerServiceMock
)

func init() {
	customerServiceMock = new(CustomerServiceMock)

	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	customer_grpc.RegisterCustomerServiceServer(s, &customerGRPCServer{
		CustomerService: customerServiceMock,
	})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCustomer(t *testing.T) {
	customer := canonical.Customer{
		Id:       "asdkjaskjd",
		Document: "123",
		Name:     "fulano",
		Email:    "fulano@mail",
		UserID:   "askdjaskdj",
	}

	customer2 := canonical.Customer{
		Id:       "asdas",
		Document: "123",
		Name:     "fulano",
		Email:    "fulano@mail",
		UserID:   "askdjaskdj",
	}

	customerServiceMock.MockGetCustomer(customer, []canonical.Customer{customer, customer2}, nil, 1)

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := customer_grpc.NewCustomerServiceClient(conn)
	resp, err := client.GetCustomer(ctx, &customer_grpc.Customer{
		Document: "123",
	})

	assert.Equal(t, customer.Id, resp.Customers[0].Id)
	assert.Equal(t, customer.Name, resp.Customers[0].Name)
	assert.Equal(t, customer.UserID, resp.Customers[0].UserId)
	assert.Equal(t, customer.Document, resp.Customers[0].Document)
	assert.Equal(t, customer.Email, resp.Customers[0].Email)

	assert.Equal(t, customer2.Id, resp.Customers[1].Id)
	assert.Equal(t, customer2.Name, resp.Customers[1].Name)
	assert.Equal(t, customer2.UserID, resp.Customers[1].UserId)
	assert.Equal(t, customer2.Document, resp.Customers[1].Document)
	assert.Equal(t, customer2.Email, resp.Customers[1].Email)

	assert.Nil(t, err)
	customerServiceMock.AssertExpectations(t)
}

func TestListenCustomer(t *testing.T) {
	go func() {
		err := ListenCustomer(7070, customerServiceMock)
		if err != nil {
			t.Errorf("Erro ao executar ListenUser: %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	serverAddr := "localhost:7070"
	_, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	assert.Nil(t, err)
	time.Sleep(100 * time.Millisecond)
}
