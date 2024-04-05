/*
Package database does all db persistence implementations.
*/
package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"literank.com/event-books/domain/model"
)

const (
	collReview = "interests"
)

// MongoPersistence runs all mongoDB operations
type MongoPersistence struct {
	db       *mongo.Database
	coll     *mongo.Collection
	pageSize int
}

// NewMongoPersistence constructs a new MongoPersistence
func NewMongoPersistence(mongoURI, dbName string, pageSize int) (*MongoPersistence, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	coll := db.Collection(collReview)
	return &MongoPersistence{db, coll, pageSize}, nil
}

// GetReview gets a review by ID
func (m *MongoPersistence) IncreaseInterest(ctx context.Context, i *model.Interest) error {
	filter := bson.M{
		"user_id": i.UserID,
		"title":   i.Title,
		"author":  i.Author,
	}
	update := bson.M{"$inc": bson.M{"score": 1}}
	opts := options.Update().SetUpsert(true)

	if _, err := m.coll.UpdateOne(ctx, filter, update, opts); err != nil {
		return err
	}
	return nil
}

// ListInterests lists user interests by a use id
func (m *MongoPersistence) ListInterests(ctx context.Context, userID string) ([]*model.Interest, error) {
	filter := bson.M{"user_id": userID}

	opts := options.Find()
	opts.SetSort(bson.M{"score": -1})
	opts.SetLimit(int64(m.pageSize))

	cursor, err := m.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	interests := make([]*model.Interest, 0)
	if err := cursor.All(ctx, &interests); err != nil {
		return nil, err
	}
	return interests, nil
}
