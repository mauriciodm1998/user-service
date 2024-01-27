package grpc

import (
	"context"
	"net"
	"strconv"
	"user-service/internal/service"
	"user-service/internal/user_grpc_files/user_grpc"

	protocol "google.golang.org/grpc"
)

type userGRPCServer struct {
	service.UserService
	user_grpc.UnimplementedUserServiceServer
}

func ListenUser(port int, service service.UserService) error {
	server := protocol.NewServer()
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	user_grpc.RegisterUserServiceServer(server, &userGRPCServer{
		UserService: service,
	})

	return server.Serve(listener)
}

func (r *userGRPCServer) GetUser(ctx context.Context, user *user_grpc.User) (*user_grpc.UserList, error) {
	request := unmarshalUser(user)

	response, err := r.UserService.GetUser(ctx, *request)
	if err != nil {
		return nil, err
	}

	return marshalUsers(response), nil
}
