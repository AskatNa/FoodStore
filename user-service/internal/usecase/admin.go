package usecase

import (
	"context"
	"fmt"
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/FoodStore/user-service/pkg/def"
	"github.com/AskatNa/FoodStore/user-service/pkg/security"
	"log"
	"time"
)

type Admin struct {
	repo       CustomerRepo
	jwtManager *security.JWTManager
}

func NewAdmin(repo CustomerRepo, jwtManager *security.JWTManager) *Admin {
	return &Admin{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (a *Admin) verifyAdminRole(token string) error {
	claims, err := a.jwtManager.Verify(token)
	if err != nil {
		return model.ErrUnauthorized
	}

	role, ok := claims["role"].(string)
	if !ok || role != model.AdminRole {
		return model.ErrForbidden
	}

	return nil
}

func (a *Admin) GetCustomerByEmail(ctx context.Context, token, email string) (model.Customer, error) {
	if err := a.verifyAdminRole(token); err != nil {
		return model.Customer{}, err
	}

	return a.repo.GetWithFilter(ctx, model.CustomerFilter{
		Email:     def.Pointer(email),
		IsDeleted: def.Pointer(false),
	})
}

func (a *Admin) UpdateCustomer(ctx context.Context, token string, email string, updates model.CustomerUpdateData) (model.Customer, error) {
	log.Printf("Admin: Attempting to update customer with email: %s", email)

	if err := a.verifyAdminRole(token); err != nil {
		log.Printf("Admin: Role verification failed: %v", err)
		return model.Customer{}, err
	}

	// First, check if customer exists
	customer, err := a.repo.GetWithFilter(ctx, model.CustomerFilter{
		Email:     def.Pointer(email),
		IsDeleted: def.Pointer(false),
	})
	if err != nil {
		log.Printf("Admin: Error finding customer: %v", err)
		return model.Customer{}, fmt.Errorf("failed to find customer: %w", err)
	}
	log.Printf("Admin: Found customer with ID: %d", customer.ID)

	// Always set UpdatedAt when updating
	updates.UpdatedAt = def.Pointer(time.Now().UTC())

	// Keep track of the email we should use to fetch the updated customer
	fetchEmail := email
	if updates.Email != nil {
		fetchEmail = *updates.Email // If email was updated, use the new email
	}

	err = a.repo.Update(ctx, model.CustomerFilter{Email: def.Pointer(email)}, updates)
	if err != nil {
		log.Printf("Admin: Error updating customer: %v", err)
		return model.Customer{}, fmt.Errorf("failed to update customer: %w", err)
	}
	log.Printf("Admin: Successfully updated customer")

	// Get updated customer data using the correct email
	updatedCustomer, err := a.repo.GetWithFilter(ctx, model.CustomerFilter{
		Email:     def.Pointer(fetchEmail),
		IsDeleted: def.Pointer(false),
	})
	if err != nil {
		log.Printf("Admin: Error getting updated customer: %v", err)
		return model.Customer{}, fmt.Errorf("failed to get updated customer: %w", err)
	}

	return updatedCustomer, nil
}

func (a *Admin) DeleteCustomer(ctx context.Context, token string, email string) error {
	if err := a.verifyAdminRole(token); err != nil {
		return err
	}

	return a.repo.Update(ctx, model.CustomerFilter{Email: def.Pointer(email)}, model.CustomerUpdateData{
		IsDeleted: def.Pointer(true),
		UpdatedAt: def.Pointer(time.Now().UTC()),
	})
}

func (a *Admin) ListCustomers(ctx context.Context, token string, filter model.CustomerFilter, pageSize, pageNumber int32) ([]model.Customer, int32, error) {
	log.Printf("Admin: Attempting to list customers with filter: %+v", filter)

	if err := a.verifyAdminRole(token); err != nil {
		log.Printf("Admin: Role verification failed: %v", err)
		return nil, 0, err
	}

	// Set IsDeleted filter to false by default if not specified
	if filter.IsDeleted == nil {
		filter.IsDeleted = def.Pointer(false)
	}

	customers, err := a.repo.GetListWithFilter(ctx, filter)
	if err != nil {
		log.Printf("Admin: Error listing customers: %v", err)
		return nil, 0, fmt.Errorf("failed to list customers: %w", err)
	}

	log.Printf("Admin: Successfully found %d customers", len(customers))
	return customers, int32(len(customers)), nil
}
