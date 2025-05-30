package frontend

import (
	"context"
	"errors"
	"github.com/AskatNa/FoodStore/user-service/internal/adapter/gRPC/frontend/dto"
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/FoodStore/user-service/pkg/security"
	base "github.com/AskatNa/apis-gen-user-service/base/frontend/v1"
	svc "github.com/AskatNa/apis-gen-user-service/service/frontend/admin/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Admin struct {
	svc.UnimplementedAdminServiceServer
	adminUseCase AdminUseCase
}

func NewAdmin(adminUseCase AdminUseCase) *Admin {
	return &Admin{
		adminUseCase: adminUseCase,
	}
}

func (a *Admin) GetCustomerByEmail(ctx context.Context, req *svc.GetCustomerByEmailRequest) (*svc.GetCustomerByEmailResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	customer, err := a.adminUseCase.GetCustomerByEmail(ctx, token, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUnauthorized):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, model.ErrForbidden):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		case errors.Is(err, model.ErrNotFound):
			return nil, status.Error(codes.NotFound, "customer not found")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &svc.GetCustomerByEmailResponse{
		Customer: dto.FromCustomerAdmin(customer),
	}, nil
}

func (a *Admin) UpdateCustomer(ctx context.Context, req *svc.UpdateCustomerRequest) (*svc.UpdateCustomerResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	updates := model.CustomerUpdateData{}
	if req.Name != nil {
		updates.Name = req.Name
	}
	if req.Phone != nil {
		updates.Phone = req.Phone
	}
	if req.NewEmail != nil {
		updates.Email = req.NewEmail
	}

	customer, err := a.adminUseCase.UpdateCustomer(ctx, token, req.Email, updates)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUnauthorized):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, model.ErrForbidden):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		case errors.Is(err, model.ErrNotFound):
			return nil, status.Error(codes.NotFound, "customer not found")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &svc.UpdateCustomerResponse{
		Customer: dto.FromCustomerAdmin(customer),
	}, nil
}

func (a *Admin) DeleteCustomer(ctx context.Context, req *svc.DeleteCustomerRequest) (*svc.DeleteCustomerResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	err := a.adminUseCase.DeleteCustomer(ctx, token, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUnauthorized):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, model.ErrForbidden):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		case errors.Is(err, model.ErrNotFound):
			return nil, status.Error(codes.NotFound, "customer not found")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &svc.DeleteCustomerResponse{}, nil
}

func (a *Admin) ListCustomers(ctx context.Context, req *svc.ListCustomersRequest) (*svc.ListCustomersResponse, error) {
	token, ok := security.TokenFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "unauthenticated")
	}

	filter := model.CustomerFilter{}
	if req.EmailFilter != nil {
		filter.Email = req.EmailFilter
	}
	if req.NameFilter != nil {
		filter.Name = req.NameFilter
	}

	customers, total, err := a.adminUseCase.ListCustomers(ctx, token, filter, req.PageSize, req.PageNumber)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrUnauthorized):
			return nil, status.Error(codes.Unauthenticated, err.Error())
		case errors.Is(err, model.ErrForbidden):
			return nil, status.Error(codes.PermissionDenied, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	customerProtos := make([]*base.Customer, len(customers))
	for i, customer := range customers {
		customerProtos[i] = dto.FromCustomerAdmin(customer)
	}

	return &svc.ListCustomersResponse{
		Customer:   customerProtos,
		TotalCount: total,
	}, nil
}
