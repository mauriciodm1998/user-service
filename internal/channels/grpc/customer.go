package grpc

import (
	"context"
	"net"
	"strconv"
	"user-service/internal/customer_grpc_files/customer_grpc"
	"user-service/internal/service"

	protocol "google.golang.org/grpc"
)

type customerGRPCServer struct {
	service.CustomerService
	customer_grpc.UnimplementedCustomerServiceServer
}

func ListenCustomer(port int, customer service.CustomerService) error {
	server := protocol.NewServer()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	customer_grpc.RegisterCustomerServiceServer(server, &customerGRPCServer{
		CustomerService: customer,
	})

	return server.Serve(listener)
}

func (r *customerGRPCServer) GetCustomer(ctx context.Context, customer *customer_grpc.Customer) (*customer_grpc.CustomerList, error) {
	request := unmarshalCustomer(customer)

	response, err := r.CustomerService.GetCustomer(ctx, *request)
	if err != nil {
		return nil, err
	}

	return marshalCustomers(response), nil
}
