package grpc

import (
	"user-service/internal/canonical"
	"user-service/internal/customer_grpc_files/customer_grpc"
	"user-service/internal/user_grpc_files/user_grpc"
)

func unmarshalCustomer(customer *customer_grpc.Customer) *canonical.Customer {
	return &canonical.Customer{
		UserID:   customer.UserId,
		Document: customer.Document,
	}
}

func marshalCustomers(customers []canonical.Customer) *customer_grpc.CustomerList {
	var response customer_grpc.CustomerList

	for _, value := range customers {
		customer := customer_grpc.Customer{
			Id:       value.Id,
			Email:    value.Email,
			Document: value.Document,
			Name:     value.Name,
			UserId:   value.UserID,
		}

		response.Customers = append(response.Customers, &customer)
	}

	return &response
}

func unmarshalUser(customer *user_grpc.User) *canonical.User {
	return &canonical.User{
		Login: customer.Login,
	}
}

func marshalUsers(users []canonical.User) *user_grpc.UserList {
	var response user_grpc.UserList

	for _, value := range users {
		user := user_grpc.User{
			Id:            value.Id,
			Login:         value.Login,
			AccessLevelID: int64(value.AccessLevelID),
		}

		response.Users = append(response.Users, &user)
	}

	return &response
}
