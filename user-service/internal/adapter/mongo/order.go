package mongo

import (
	"context"
	"github.com/AskatNa/FoodStore/user-service/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		collection: db.Collection("orders"),
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return err
	}
	order.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *OrderRepository) GetOrder(ctx context.Context, id primitive.ObjectID) (*model.Order, error) {
	var order model.Order
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) ListOrders(ctx context.Context, customerID primitive.ObjectID, pageSize, pageNumber int32, status string, fromDate, toDate time.Time) ([]*model.Order, int32, error) {
	filter := bson.M{"customer_id": customerID}

	if status != "" {
		filter["status"] = status
	}

	if !fromDate.IsZero() || !toDate.IsZero() {
		dateFilter := bson.M{}
		if !fromDate.IsZero() {
			dateFilter["$gte"] = fromDate
		}
		if !toDate.IsZero() {
			dateFilter["$lte"] = toDate
		}
		filter["created_at"] = dateFilter
	}

	opts := options.Find().
		SetSkip(int64((pageNumber - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var orders []*model.Order
	if err = cursor.All(ctx, &orders); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return orders, int32(total), nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
