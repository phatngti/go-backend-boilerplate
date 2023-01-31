package core_crud

import (
	"context"
	"fmt"
	"log"
	"phatngti/boilerplate/core"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository[E any] struct {
	collection *mongo.Collection
}

func NewMongoRepository[E any](collection *mongo.Collection) *MongoRepository[E] {
	return &MongoRepository[E]{
		collection: collection,
	}
}

func (m MongoRepository[E]) GetAll(ctx context.Context, criteria map[string]interface{}) ([]E, error) {
	filter := core.MapToBson(criteria)
	cursor, err := m.collection.Find(ctx, filter)
	if err != nil {
		log.Fatal("Get all error: ", err)
		return *new([]E), nil
	}

	var result []bson.D
	if err = cursor.All(ctx, &result); err != nil {
		log.Fatal("Get all failed", err)
		return *new([]E), err
	}
	fmt.Println("result: ", result)
	return *new([]E), nil
}

// func (m MongoRepository[E]) Insert(ctx context.Context, entity *E) (E, error) {
// 	doc := core.StructToBSON(entity)
// 	fmt.Println("doc: ", doc)
// 	rs, err := m.collection.InsertOne(ctx, doc)
// 	if err != nil {
// 		return *new(E), err
// 	}
// 	entity.ID = rs.InsertedID
// 	return *entity, nil
// }
