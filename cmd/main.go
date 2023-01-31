package main

import (
	"fmt"
	"log"
	"os"
	"phatngti/boilerplate/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func initDatabase() *database.Database {
	database := new(database.Database)
	if _, err := strconv.ParseBool(os.Getenv("RUN_PSQL")); err == nil {
		dns := fmt.Sprintf("host=%s  port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"), os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"), os.Getenv("PSQL_DB"))
		database.InitPSql(dns)
	}
	return database
}
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error: Failed to load env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init database and models
	// database := initDatabase()
	// fmt.Println("database: ", database)
	// if database.GetPSqlDB() != nil {
	// 	models := new(models.Models)
	// 	models.InitModels(database)
	// }
	var compoundIndex bson.D
	fmt.Println("compoundIndex: ", compoundIndex)

	fmt.Println("print bson: ", bson.D{{Key: "test", Value: 123}, {Key: "abc", Value: 456}})
	fmt.Println("print bson: ", bson.D{{Key: "test", Value: 123}, {Key: "abc", Value: 456}})

	// // Start default gin server
	// server := gin.Default()

	// server.Use(gzip.Gzip(gzip.BestCompression))

}
