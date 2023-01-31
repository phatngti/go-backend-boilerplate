package core_crud_test

import (
	"context"
	core_crud "phatngti/boilerplate/core/crud"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string              `bson:"name,omitempty" json:"name,omitempty"`
	Weight    uint                `bson:"weight,omitempty" json:"weight,omitempty"`
	CreatedAt *time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time          `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

func ProductRepository(db *mongo.Database) *core_crud.MongoRepository[Product] {
	collection := db.Collection("products")
	productRepo := core_crud.NewMongoRepository[Product](collection)
	return productRepo
}

func GetDB() *mongo.Database {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://phatnt:rE1MWr8ngUpB9sVB@cluster0.upbtgmz.mongodb.net/test"))
	if err != nil {
		panic(err)
	}
	database := client.Database("test")
	return database
}

// func TestMongoGetAll(t *testing.T) {
// 	db := GetDB()
// 	productRepo := ProductRepository(db)

// }

// func TestInsertOne(t *testing.T) {
// 	db := GetDB()
// 	ctx := context.Background()
// 	productRepo := ProductRepository(db)
// 	now := time.Now().UTC()
// 	doc := &Product{
// 		Name:      "abc",
// 		CreatedAt: &now,
// 	}
// 	fmt.Println("doc: ", doc)
// 	productRepo.Insert(ctx, doc)
// }
