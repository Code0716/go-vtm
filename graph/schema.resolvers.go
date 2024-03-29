package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/Code0716/go-vtm/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	newUser, err := r.Handler.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// CreateAddress is the resolver for the createAddress field.
func (r *mutationResolver) CreateAddress(ctx context.Context, input model.AddressInput) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: CreateAddress - createAddress"))
}

// UpdateAddress is the resolver for the updateAddress field.
func (r *mutationResolver) UpdateAddress(ctx context.Context, input model.AddressInput) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: UpdateAddress - updateAddress"))
}

// DeleteAddress is the resolver for the deleteAddress field.
func (r *mutationResolver) DeleteAddress(ctx context.Context, id string) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: DeleteAddress - deleteAddress"))
}

// CreateDepartment is the resolver for the createDepartment field.
func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.DepartmentInput) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: CreateDepartment - createDepartment"))
}

// UpdateDepartment is the resolver for the updateDepartment field.
func (r *mutationResolver) UpdateDepartment(ctx context.Context, input model.DepartmentInput) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: UpdateDepartment - updateDepartment"))
}

// DeleteDepartment is the resolver for the deleteDepartment field.
func (r *mutationResolver) DeleteDepartment(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: DeleteDepartment - deleteDepartment"))
}

// CreateAttendance is the resolver for the createAttendance field.
func (r *mutationResolver) CreateAttendance(ctx context.Context, input model.AttendanceInput) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: CreateAttendance - createAttendance"))
}

// UpdateAttendance is the resolver for the updateAttendance field.
func (r *mutationResolver) UpdateAttendance(ctx context.Context, input model.AttendanceInput) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: UpdateAttendance - updateAttendance"))
}

// DeleteAttendance is the resolver for the deleteAttendance field.
func (r *mutationResolver) DeleteAttendance(ctx context.Context, id string) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: DeleteAttendance - deleteAttendance"))
}

// CreateInvoice is the resolver for the createInvoice field.
func (r *mutationResolver) CreateInvoice(ctx context.Context, input model.InvoiceInput) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: CreateInvoice - createInvoice"))
}

// UpdateInvoice is the resolver for the updateInvoice field.
func (r *mutationResolver) UpdateInvoice(ctx context.Context, input model.InvoiceInput) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: UpdateInvoice - updateInvoice"))
}

// DeleteInvoice is the resolver for the deleteInvoice field.
func (r *mutationResolver) DeleteInvoice(ctx context.Context, id string) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: DeleteInvoice - deleteInvoice"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// ListUsers is the resolver for the listUsers field.
func (r *queryResolver) ListUsers(ctx context.Context, input *model.LimitOffset) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: ListUsers - listUsers"))
}

// GetAddress is the resolver for the getAddress field.
func (r *queryResolver) GetAddress(ctx context.Context, id string) (*model.Address, error) {
	panic(fmt.Errorf("not implemented: GetAddress - getAddress"))
}

// ListAddress is the resolver for the listAddress field.
func (r *queryResolver) ListAddress(ctx context.Context, input *model.LimitOffset) ([]*model.Address, error) {
	panic(fmt.Errorf("not implemented: ListAddress - listAddress"))
}

// GetDepartment is the resolver for the getDepartment field.
func (r *queryResolver) GetDepartment(ctx context.Context, id string) (*model.Department, error) {
	panic(fmt.Errorf("not implemented: GetDepartment - getDepartment"))
}

// ListDepartment is the resolver for the listDepartment field.
func (r *queryResolver) ListDepartment(ctx context.Context, input *model.LimitOffset) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented: ListDepartment - listDepartment"))
}

// GetAttendance is the resolver for the getAttendance field.
func (r *queryResolver) GetAttendance(ctx context.Context, id string) (*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: GetAttendance - getAttendance"))
}

// ListAttendance is the resolver for the listAttendance field.
func (r *queryResolver) ListAttendance(ctx context.Context, input *model.ListAttendanceInput) ([]*model.Attendance, error) {
	panic(fmt.Errorf("not implemented: ListAttendance - listAttendance"))
}

// GetInvoice is the resolver for the getInvoice field.
func (r *queryResolver) GetInvoice(ctx context.Context, userID string) (*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: GetInvoice - getInvoice"))
}

// ListInvoice is the resolver for the listInvoice field.
func (r *queryResolver) ListInvoice(ctx context.Context, input *model.ListInvoiceInput) ([]*model.Invoice, error) {
	panic(fmt.Errorf("not implemented: ListInvoice - listInvoice"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
