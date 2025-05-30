package mongo

import (
	"context"
	"github.com/AskatNa/FoodStore/user-service/internal/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) ListProducts(ctx context.Context, pageSize, pageNumber int32, category string) ([]*model.Product, int32, error) {
	filter := bson.M{}
	if category != "" {
		filter["category"] = category
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

	var products []*model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return products, int32(total), nil
}

func (r *ProductRepository) GetProduct(ctx context.Context, id primitive.ObjectID) (*model.Product, error) {
	var product model.Product
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) SearchProducts(ctx context.Context, query string, pageSize, pageNumber int32) ([]*model.Product, int32, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"description": bson.M{"$regex": query, "$options": "i"}},
			{"category": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	opts := options.Find().
		SetSkip(int64((pageNumber - 1) * pageSize)).
		SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []*model.Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return products, int32(total), nil
}
