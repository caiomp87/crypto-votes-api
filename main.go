package main

import (
	"fmt"
	"log"
	"os"

	"github.com/caiomp87/crypto-votes-api/routes"
	"github.com/caiomp87/crypto-votes-api/utils"
	"github.com/caiomp87/crypto-votes-datastore/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	connectionString, err := utils.BuildConnectionString()
	if err != nil {
		log.Fatalf("could not get build connection string: %v", err)
	}

	postgresService := postgres.NewPostresService("postgres", connectionString)

	db, err := postgresService.Connect()
	if err != nil {
		log.Fatalf("could not connect on database: %v", err)
	}

	defer func() {
		if err := postgresService.Disconnect(db); err != nil {
			log.Fatalf("could not disconnect from database gracefullly: %v", err)
		}
	}()

	postgres.CryptoDatastore = postgres.NewCryptoService(db)

	env, ok := os.LookupEnv("ENV")
	if env == "production" || !ok {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routes.AddCryptoRoutes(r)

	fmt.Println("server listening on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
