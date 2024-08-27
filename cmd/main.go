package main

import (
	"database/sql"
	"github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters"
	"github.com/Pr3c10us/boilerplate/internals/infrastructures/ports"
	"github.com/Pr3c10us/boilerplate/internals/services"
	"github.com/Pr3c10us/boilerplate/packages/configs"
	"github.com/Pr3c10us/boilerplate/packages/logger"
	"github.com/Pr3c10us/boilerplate/packages/utils"
	"github.com/redis/go-redis/v9"
	"github.com/stripe/stripe-go/v79"
)

var (
	environmentVariables = configs.LoadEnvironment()
)

func init() {
	configs.Goth(environmentVariables)
}

func main() {
	newLogger := logger.NewSugarLogger(environmentVariables.ProductionEnvironment)
	newPGConnection := utils.NewPGConnection(environmentVariables)
	defer func(db *sql.DB) {
		_ = db.Close()
	}(newPGConnection)
	newS3Client := utils.NewS3Client(environmentVariables)
	newRedisConnection := utils.NewRedisClient(environmentVariables)
	defer func(redis *redis.Client) {
		_ = redis.Close()
	}(newRedisConnection)
	newSESClient := utils.NewSESClient(environmentVariables)
	newSNSClient := utils.NewSNSClient(environmentVariables)
	stripe.Key = environmentVariables.Stripe.SecretKey

	adapterDependencies := adapters.AdapterDependencies{
		Logger:               newLogger,
		EnvironmentVariables: environmentVariables,
		DB:                   newPGConnection,
		Redis:                newRedisConnection,
		S3Client:             newS3Client,
		SESClient:            newSESClient,
		SNSClient:            newSNSClient,
	}
	newAdapters := adapters.NewAdapters(adapterDependencies)
	newServices := services.NewServices(newAdapters)
	newPort := ports.NewPorts(newServices, newLogger, environmentVariables)
	newPort.GinServer.Run()
}
