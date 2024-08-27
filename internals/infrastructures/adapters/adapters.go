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
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/redis/go-redis/v9"
)

type AdapterDependencies struct {
	Logger               logger.Logger
	EnvironmentVariables *configs.EnvironmentVariables
	DB                   *sql.DB
	Redis                *redis.Client
	S3Client             storage.S3ClientInterface
	SESClient            *ses.Client
	SNSClient            *sns.Client
}

type Adapters struct {
	Logger                   logger.Logger
	EnvironmentVariables     *configs.EnvironmentVariables
	AuthenticationRepository authentication.Repository
	EmailRepository          email.Repository
	SMSRepository            sms.Repository
}

func NewAdapters(dependencies AdapterDependencies) *Adapters {
	return &Adapters{
		Logger:                   dependencies.Logger,
		EnvironmentVariables:     dependencies.EnvironmentVariables,
		AuthenticationRepository: authentication2.NewAuthenticationRepositoryPG(dependencies.DB),
		EmailRepository:          email2.NewGoMailEmailRepository(dependencies.EnvironmentVariables),
	}
}
