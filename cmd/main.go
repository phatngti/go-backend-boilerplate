package cmd

import (
	"log"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error: Failed to load env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Start default gin server
	server := gin.Default()

	server.Use(gzip.Gzip(gzip.BestCompression))


}
