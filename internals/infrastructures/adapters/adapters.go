package adapters

import (
	"database/sql"
	"github.com/Pr3c10us/boilerplate/internals/domains/authentication"
	"github.com/Pr3c10us/boilerplate/internals/domains/email"
	"github.com/Pr3c10us/boilerplate/internals/domains/sms"
	"github.com/Pr3c10us/boilerplate/internals/domains/storage"
	authentication2 "github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters/authentication"
	email2 "github.com/Pr3c10us/boilerplate/internals/infrastructures/adapters/email"
	"github.com/Pr3c10us/boilerplate/packages/configs"
	"github.com/Pr3c10us/boilerplate/packages/logger"
	"github.com/redis/go-redis/v9"
)

type Adapters struct {
	Logger                   logger.Logger
	EnvironmentVariables     *configs.EnvironmentVariables
	AuthenticationRepository authentication.Repository
	EmailRepository          email.Repository
	SMSRepository            sms.Repository
}

func NewAdapters(logger logger.Logger, environmentVariables *configs.EnvironmentVariables, db *sql.DB, redis *redis.Client, s3Client storage.S3ClientInterface) *Adapters {
	return &Adapters{
		Logger:                   logger,
		EnvironmentVariables:     environmentVariables,
		AuthenticationRepository: authentication2.NewAuthenticationRepositoryPG(db),
		EmailRepository:          email2.NewGoMailEmailRepository(environmentVariables),
	}
}
