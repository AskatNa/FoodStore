package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/AskatNa/FoodStore/user-service/internal/adapter/mongo/dao"
	"github.com/AskatNa/FoodStore/user-service/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Customer struct {
	conn       *mongo.Database
	collection string
}

const (
	collectionCustomers = "customers"
)

func NewCustomer(conn *mongo.Database) *Customer {
	return &Customer{
		conn:       conn,
		collection: collectionCustomers,
	}
}

func (a *Customer) EnsureIndexes(ctx context.Context) error {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"email": 1},
		Options: options.Index().
			SetUnique(true).
			SetName("unique_email"),
	}

	_, err := a.conn.Collection(a.collection).Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return fmt.Errorf("failed to create unique index on email: %w", err)
	}

	return nil
}

func (a *Customer) Create(ctx context.Context, customer model.Customer) error {
	log.Printf("MongoDB: Attempting to create customer with ID: %d, Email: %s", customer.ID, customer.Email)

	daoCustomer := dao.FromCustomer(customer)
	log.Printf("MongoDB: Converted to DAO model successfully")

	_, err := a.conn.Collection(a.collection).InsertOne(ctx, daoCustomer)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Printf("MongoDB: Duplicate key error - email already exists: %s", customer.Email)
			return model.ErrEmailAlreadyRegistered
		}
		log.Printf("MongoDB: Error creating customer: %v", err)
		return fmt.Errorf("customer with ID %d has not been created: %w", customer.ID, err)
	}

	log.Printf("MongoDB: Successfully created customer with ID: %d", customer.ID)
	return nil
}

func (a *Customer) Update(ctx context.Context, filter model.CustomerFilter, update model.CustomerUpdateData) error {
	log.Printf("MongoDB Update - Filter: %+v", filter)
	mongoFilter := dao.FromCustomerFilter(filter)
	log.Printf("MongoDB Update - MongoDB Filter: %+v", mongoFilter)

	updateDoc := dao.FromCustomerUpdateData(update)
	log.Printf("MongoDB Update - Update Document: %+v", updateDoc)

	result, err := a.conn.Collection(a.collection).UpdateOne(
		ctx,
		mongoFilter,
		updateDoc,
	)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return model.ErrEmailAlreadyRegistered
		}
		return fmt.Errorf("customer has not been updated with filter: %v, err: %w", filter, err)
	}

	log.Printf("MongoDB Update - Modified %d documents", result.ModifiedCount)
	if result.ModifiedCount == 0 {
		log.Printf("MongoDB Update - No documents were modified!")
	}

	return nil
}

func (a *Customer) GetWithFilter(ctx context.Context, filter model.CustomerFilter) (model.Customer, error) {
	log.Printf("MongoDB GetWithFilter - Filter: %+v", filter)
	mongoFilter := dao.FromCustomerFilter(filter)
	log.Printf("MongoDB GetWithFilter - MongoDB Filter: %+v", mongoFilter)

	var customerDAO dao.Customer
	err := a.conn.Collection(a.collection).FindOne(ctx, mongoFilter).Decode(&customerDAO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("MongoDB GetWithFilter - No document found for filter")
			return model.Customer{}, model.ErrNotFound
		}
		log.Printf("MongoDB GetWithFilter - Error: %v", err)
		return model.Customer{}, fmt.Errorf("collection.FindOne: %w", err)
	}

	log.Printf("MongoDB GetWithFilter - Found customer: ID=%d, Email=%s", customerDAO.ID, customerDAO.Email)
	return dao.ToCustomer(customerDAO), nil
}

func (a *Customer) GetListWithFilter(ctx context.Context, filter model.CustomerFilter) ([]model.Customer, error) {
	log.Printf("MongoDB: Listing all customers in database...")

	// First, let's see all customers without any filter
	cursor, err := a.conn.Collection(a.collection).Find(ctx, bson.M{})
	if err != nil {
		log.Printf("MongoDB: Error listing customers: %v", err)
		return nil, fmt.Errorf("collection.Find: %w", err)
	}
	defer cursor.Close(ctx)

	var customersDAO []dao.Customer
	if err = cursor.All(ctx, &customersDAO); err != nil {
		log.Printf("MongoDB: Error decoding customers: %v", err)
		return nil, fmt.Errorf("cursor.All: %w", err)
	}

	log.Printf("MongoDB: === ALL CUSTOMERS IN DATABASE ===")
	log.Printf("MongoDB: Found %d total customers", len(customersDAO))
	for _, c := range customersDAO {
		log.Printf("MongoDB: Customer - ID: %d, Email: %s, Name: %s, IsDeleted: %v, CreatedAt: %v",
			c.ID, c.Email, c.Name, c.IsDeleted, c.CreatedAt)
	}
	log.Printf("MongoDB: === END OF CUSTOMER LIST ===")

	// Now apply the filter if any
	filterDoc := dao.FromCustomerFilter(filter)
	log.Printf("MongoDB: Applying filter: %+v", filterDoc)

	cursor, err = a.conn.Collection(a.collection).Find(ctx, filterDoc)
	if err != nil {
		return nil, fmt.Errorf("collection.Find: %w", err)
	}
	defer cursor.Close(ctx)

	customersDAO = nil
	if err = cursor.All(ctx, &customersDAO); err != nil {
		return nil, fmt.Errorf("cursor.All: %w", err)
	}

	customers := make([]model.Customer, len(customersDAO))
	for i, customerDAO := range customersDAO {
		customers[i] = dao.ToCustomer(customerDAO)
	}

	log.Printf("MongoDB: Returning %d customers after applying filter", len(customers))
	return customers, nil
}

func (a *Customer) HardDelete(ctx context.Context, filter model.CustomerFilter) error {
	_, err := a.conn.Collection(a.collection).DeleteOne(ctx, dao.FromCustomerFilter(filter))
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}
	return nil
}
