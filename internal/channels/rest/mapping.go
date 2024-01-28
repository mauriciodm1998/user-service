package rest

import "user-service/internal/canonical"

func (c *CreateUserRequest) toCanonical() (canonical.Customer, canonical.User) {
	return canonical.Customer{
			Name:     c.Name,
			Document: c.Document,
			Email:    c.Email,
		}, canonical.User{
			Login:    c.Login,
			Password: c.Password,
		}
}

func customerToResponse(customer canonical.Customer) CustomerResponse {
	return CustomerResponse{
		Id:       customer.Id,
		UserID:   customer.UserID,
		Document: customer.Document,
		Name:     customer.Name,
		Email:    customer.Email,
	}
}

func userToResponse(user canonical.User) UserResponse {
	return UserResponse{
		Id:            user.Id,
		Login:         user.Login,
		AccessLevelID: user.AccessLevelID,
	}
}
