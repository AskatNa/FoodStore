package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/FoodStore/user-service/pkg/def"
	"github.com/AskatNa/FoodStore/user-service/pkg/security"
	"github.com/AskatNa/FoodStore/user-service/pkg/transactor"
	"log"
	"time"
)

type Customer struct {
	ai              AiRepo
	repo            CustomerRepo
	tokenRepo       RefreshTokenRepo
	producer        CustomerEventStorage
	callTx          transactor.WithinTransactionFunc
	jwtManager      *security.JWTManager
	passwordManager *security.PasswordManager
}

func NewCustomer(
	ai AiRepo,
	repo CustomerRepo,
	tokenRepo RefreshTokenRepo,
	producer CustomerEventStorage,
	callTx transactor.WithinTransactionFunc,
	jwtManager *security.JWTManager,
	passwordManager *security.PasswordManager,
) *Customer {
	return &Customer{
		ai:              ai,
		repo:            repo,
		tokenRepo:       tokenRepo,
		producer:        producer,
		callTx:          callTx,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}

func (uc *Customer) Register(ctx context.Context, request model.Customer) (uint64, error) {
	log.Printf("Starting registration for email: %s", request.Email)

	// Generate ID
	id, err := uc.ai.Next(ctx, model.CustomerAi)
	if err != nil {
		log.Printf("Error getting next ID: %v", err)
		return 0, fmt.Errorf("ai.Next: %w", err)
	}
	log.Printf("Generated ID: %d", id)
	request.ID = id

	// Hash password
	request.PasswordHash, err = uc.passwordManager.HashPassword(request.NewPassword)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return 0, fmt.Errorf("passwordManager.HashPassword: %w", err)
	}
	log.Printf("Password hashed successfully")

	// Set timestamps
	request.CreatedAt = time.Now().UTC()
	request.UpdatedAt = time.Now().UTC()

	// Create customer
	log.Printf("Attempting to create customer in database")
	err = uc.repo.Create(ctx, request)
	if err != nil {
		log.Printf("Error creating customer: %v", err)
		return 0, fmt.Errorf("repo.Create: %w", err)
	}
	log.Printf("Customer created successfully")

	// Push event (non-critical operation)
	log.Printf("Attempting to push customer event")
	if err := uc.producer.Push(ctx, request); err != nil {
		log.Printf("Warning: Error pushing customer event: %v", err)
		// Don't return error here as this is non-critical
	}

	log.Printf("Registration completed successfully with ID: %d", request.ID)
	return request.ID, nil
}

func (uc *Customer) Update(ctx context.Context, token string, request model.Customer) (model.Customer, error) {
	claims, err := uc.jwtManager.Verify(token)
	if err != nil {
		return model.Customer{}, model.ErrInvalidID
	}
	customerID, ok := claims["user_id"].(float64)
	if !ok {
		return model.Customer{}, model.ErrInvalidID
	}
	if uint64(customerID) != request.ID {
		return model.Customer{}, model.ErrInvalidID
	}

	dbCustomer, err := uc.Get(ctx, token, request.ID)
	if err != nil {
		return model.Customer{}, err
	}

	// Only verify password if a new password is being set
	if request.NewPassword != "" {
		if request.CurrentPassword == "" {
			return model.Customer{}, fmt.Errorf("current password is required when updating password")
		}
		err = uc.passwordManager.CheckPassword(dbCustomer.PasswordHash, request.CurrentPassword)
		if err != nil {
			return model.Customer{}, fmt.Errorf("uc.passwordManager.CheckPassword: %w", err)
		}
		request.PasswordHash, err = uc.passwordManager.HashPassword(request.NewPassword)
		if err != nil {
			return model.Customer{}, fmt.Errorf("uc.passwordManager.HashPassword: %w", err)
		}
	}

	// Build update data with only provided fields
	updateData := model.CustomerUpdateData{
		ID:        def.Pointer(request.ID),
		UpdatedAt: def.Pointer(time.Now().UTC()),
	}

	if request.Name != "" {
		updateData.Name = def.Pointer(request.Name)
	}
	if request.Phone != "" {
		updateData.Phone = def.Pointer(request.Phone)
	}
	if request.Email != "" {
		updateData.Email = def.Pointer(request.Email)
	}
	if request.PasswordHash != "" {
		updateData.PasswordHash = def.Pointer(request.PasswordHash)
	}

	err = uc.repo.Update(ctx, model.CustomerFilter{ID: &request.ID}, updateData)
	if err != nil {
		return model.Customer{}, err
	}

	// Get updated customer data
	updatedCustomer, err := uc.Get(ctx, token, request.ID)
	if err != nil {
		return model.Customer{}, err
	}

	err = uc.producer.Push(ctx, updatedCustomer)
	if err != nil {
		log.Println("uc.producer.Push: %w", err)
	}

	return updatedCustomer, nil
}

func (uc *Customer) Get(ctx context.Context, token string, id uint64) (model.Customer, error) {
	claims, err := uc.jwtManager.Verify(token)
	if err != nil {
		return model.Customer{}, model.ErrInvalidID
	}

	customerID, ok := claims["user_id"].(float64)
	if !ok {
		return model.Customer{}, model.ErrInvalidID
	}

	if uint64(customerID) != id {
		return model.Customer{}, model.ErrInvalidID
	}

	return uc.repo.GetWithFilter(ctx, model.CustomerFilter{ID: &id})
}

func (uc *Customer) Delete(ctx context.Context, id uint64) error {
	if id <= 0 {
		return model.ErrInvalidID
	}

	err := uc.repo.Update(ctx, model.CustomerFilter{ID: &id}, model.CustomerUpdateData{
		IsDeleted: def.Pointer(true),
		UpdatedAt: def.Pointer(time.Now().UTC()),
	})
	if err != nil {
		return fmt.Errorf("uc.repo.Update: %w", err)
	}

	return nil
}

func (uc *Customer) Login(ctx context.Context, email, password string) (model.Token, error) {
	// Check for admin credentials first
	log.Printf("Login attempt - email: %s", email)

	if email == model.AdminEmail && password == model.AdminPassword {
		log.Printf("Admin credentials match - generating tokens")
		// Generate admin tokens
		accessToken, err := uc.jwtManager.GenerateAccessToken(1, model.AdminRole) // Using ID 1 for admin
		if err != nil {
			log.Printf("Error generating admin access token: %v", err)
			return model.Token{}, err
		}
		refreshToken, err := uc.jwtManager.GenerateRefreshToken(1)
		if err != nil {
			log.Printf("Error generating admin refresh token: %v", err)
			return model.Token{}, err
		}

		session := model.Session{
			UserID:       1, // Admin ID is 1
			RefreshToken: refreshToken,
			ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
			CreatedAt:    time.Now(),
		}

		err = uc.tokenRepo.Create(ctx, session)
		if err != nil {
			log.Printf("Error creating admin session: %v", err)
			return model.Token{}, err
		}

		log.Printf("Admin login successful")
		return model.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}

	log.Printf("Not admin credentials, checking regular user")
	// If not admin credentials, proceed with regular customer login
	customer, err := uc.repo.GetWithFilter(ctx, model.CustomerFilter{Email: def.Pointer(email)})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			log.Printf("User not found with email: %s", email)
			return model.Token{}, model.ErrNotFound
		}
		log.Printf("Error getting customer: %v", err)
		return model.Token{}, fmt.Errorf("failed to get customer: %w", err)
	}

	err = uc.passwordManager.CheckPassword(customer.PasswordHash, password)
	if err != nil {
		log.Printf("Invalid password for user: %s", email)
		return model.Token{}, model.ErrNotFound // Return not found for wrong password too for security
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(customer.ID, model.CustomerRole)
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		return model.Token{}, fmt.Errorf("failed to generate access token: %w", err)
	}
	refreshToken, err := uc.jwtManager.GenerateRefreshToken(customer.ID)
	if err != nil {
		log.Printf("Error generating refresh token: %v", err)
		return model.Token{}, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	session := model.Session{
		UserID:       customer.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, session)
	if err != nil {
		log.Printf("Error creating session: %v", err)
		return model.Token{}, fmt.Errorf("failed to create session: %w", err)
	}

	log.Printf("Regular user login successful - ID: %d", customer.ID)
	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *Customer) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	session, err := uc.tokenRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}
	if session.ExpiresAt.Before(time.Now().UTC()) {
		return model.Token{}, model.ErrRefreshTokenExpired
	}

	customer, err := uc.repo.GetWithFilter(ctx, model.CustomerFilter{ID: def.Pointer(session.UserID)})
	if err != nil {
		return model.Token{}, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(customer.ID, model.CustomerRole)
	if err != nil {
		return model.Token{}, err
	}

	newRefreshToken, err := uc.jwtManager.GenerateRefreshToken(customer.ID)
	if err != nil {
		return model.Token{}, err
	}

	// delete old refresh and insert new one (rotation)
	err = uc.tokenRepo.DeleteByToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}

	newSession := model.Session{
		UserID:       customer.ID,
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, newSession)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
