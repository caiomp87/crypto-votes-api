package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getDBHost() string {
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		dbHost = "localhost"
	}

	return dbHost
}

func getDBPort() int {
	dbPort := 5432

	dbPortStr, ok := os.LookupEnv("DB_PORT")
	if ok {
		convertedDbPort, err := strconv.Atoi(dbPortStr)
		if err == nil && convertedDbPort != 0 {
			dbPort = convertedDbPort
		}
	}

	return dbPort
}

func getDBuser() (string, error) {
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		return "", errors.New(`environment variable "DB_USER" is required`)
	}

	return dbUser, nil
}

func getDBPassword() (string, error) {
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		return "", errors.New(`environment variable "DB_PASSWORD" is required`)
	}

	return dbPassword, nil
}

func getDBName() (string, error) {
	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return "", errors.New(`environment variable "DB_NAME" is required`)
	}

	return dbName, nil
}

func BuildConnectionString() (string, error) {
	dbUser, err := getDBuser()
	if err != nil {
		log.Fatalf("could not get database user: %v", err)
	}

	dbPassword, err := getDBPassword()
	if err != nil {
		log.Fatalf("could not get database password: %v", err)
	}

	dbName, err := getDBName()
	if err != nil {
		log.Fatalf("could not get database name: %v", err)
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbUser, dbPassword, getDBHost(), getDBPort(), dbName)

	return connectionString, nil
}
