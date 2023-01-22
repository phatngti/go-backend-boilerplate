package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	postgresDb *gorm.DB
	mongoDb    *mongo.Client
}

func (x *Database) InitPSql(dns string) {
	db, err := connectPSqlDB(dns)
	if err != nil {
		panic(err)
	}
	x.postgresDb = db
	return
}

func (x *Database) InitMongo(dns string) {
	client, err := connectMongoDB(dns)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	x.mongoDb = client
	return
}

func connectPSqlDB(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect postgres db")
		return nil, err
	}
	return db, nil
}

func connectMongoDB(dataSourceName string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(dataSourceName))
	if err != nil {
		log.Fatal("Failed to connect database")
		return nil, err
	}

	return client, nil
}

func (x Database) GetPSqlDB() *gorm.DB {
	return x.postgresDb
}

func (x Database) GetMongoDB() *mongo.Client {
	return x.mongoDb
}
