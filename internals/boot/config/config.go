package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	// Port application running port
	Port string = ":"

	// DatabaseHost provides the host for the database
	DatabaseHost string

	// DatabaseName provides name of the database
	DatabaseName string

	// JWTAccessTokenTime for access token
	JWTAccessTokenTime time.Duration

	// JWTRefreshTokenTime for access token
	JWTRefreshTokenTime time.Duration

	// TokenSignature token signature string
	TokenSignature string

	// SendGridAPIKey
	SendGridAPIKey string
)

// getEnvValue gets value from .env file
func getEnvValue(key string) interface{} {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return string(os.Getenv(key))
}

// Init initialize configuration
func Init() {
	strToInt := func(s interface{}) int64 {
		strInt, _ := s.(int64)
		return strInt
	}

	Port = ":" + getEnvValue("PORT").(string)
	DatabaseHost = getEnvValue("DATABASE_HOST").(string)
	DatabaseName = getEnvValue("DATABASE_NAME").(string)
	JWTAccessTokenTime = (time.Minute * time.Duration(strToInt(getEnvValue("ACCESS_TOKEN_TIMING"))))
	JWTRefreshTokenTime = (time.Minute * time.Duration(strToInt(getEnvValue("REFRESH_TOKEN_TIMING"))))
	TokenSignature = getEnvValue("TOKEN_SIGNATURE").(string)
	SendGridAPIKey = getEnvValue("SENDGRID_API_KEY").(string)
}
