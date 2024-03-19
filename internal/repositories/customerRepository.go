package repository

import (
	"context"
	"user-service/internal/canonical"
)

type CustomerRepository interface {
	CreateCustomer(context.Context, canonical.Customer) error
	GetCustomerByUserId(context.Context, string) (*canonical.Customer, error)
	GetCustomerByDocument(context.Context, string) (*canonical.Customer, error)
	GetAllCustomers(ctx context.Context) ([]canonical.Customer, error)
	DeleteCustomer(ctx context.Context, customerId string) error
}

type customerRepository struct {
	db PgxIface
}

func NewCustomerRepo(db PgxIface) CustomerRepository {
	return &customerRepository{db}
}

func (r *customerRepository) CreateCustomer(ctx context.Context, customer canonical.Customer) error {
	sqlStatement := "INSERT INTO \"Customer\" (Id, UserID, Document, Name, Email) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.db.Exec(ctx, sqlStatement, customer.Id, customer.UserID, customer.Document, customer.Name, customer.Email)
	if err != nil {
		return err
	}
	return nil
}

func (r *customerRepository) GetCustomerByUserId(ctx context.Context, userID string) (*canonical.Customer, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Customer\" WHERE UserID = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customer canonical.Customer
	if rows.Next() {
		if err = rows.Scan(
			&customer.Id,
			&customer.UserID,
			&customer.Document,
			&customer.Name,
			&customer.Email,
		); err != nil {
			return nil, err
		}
		return &customer, nil
	}

	return nil, ErrorNotFound
}

func (r *customerRepository) GetCustomerByDocument(ctx context.Context, document string) (*canonical.Customer, error) {
	rows, err := r.db.Query(ctx,
		"SELECT * FROM \"Customer\" WHERE Document = $1",
		document,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customer canonical.Customer
	if rows.Next() {
		if err = rows.Scan(
			&customer.Id,
			&customer.UserID,
			&customer.Document,
			&customer.Name,
			&customer.Email,
		); err != nil {
			return nil, err
		}
		return &customer, nil
	}

	return nil, ErrorNotFound
}

func (r *customerRepository) GetAllCustomers(ctx context.Context) ([]canonical.Customer, error) {
	rows, err := r.db.Query(ctx, "SELECT * FROM \"Customer\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []canonical.Customer

	for rows.Next() {
		var customer canonical.Customer

		if err = rows.Scan(
			&customer.Id,
			&customer.UserID,
			&customer.Document,
			&customer.Name,
			&customer.Email,
		); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (r *customerRepository) DeleteCustomer(ctx context.Context, customerId string) error {
	_, err := r.db.Exec(ctx, "DELETE FROM  \"Customer\" WHERE ID = $1", customerId)
	if err != nil {
		return err
	}

	return nil
}
