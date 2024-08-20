package utils

import (
	"database/sql"
	"fmt"
	"github.com/Pr3c10us/boilerplate/packages/configs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	redis "github.com/redis/go-redis/v9"
)

func NewPGConnection(env *configs.EnvironmentVariables) *sql.DB {
	// PG_DB instantiation
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		env.PostgresDB.Host,
		env.PostgresDB.Port,
		env.PostgresDB.Username,
		env.PostgresDB.Password,
		env.PostgresDB.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("failed to instantiation DB connection")
	}
	err = db.Ping()
	if err != nil {
		panic("no connection could be made because the target machine actively refused it")
	}
	return db
}

func NewRedisClient(env *configs.EnvironmentVariables) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     env.RedisCache.Address,
		Password: env.RedisCache.Password,
		Username: env.RedisCache.Username,
		DB:       0,
	})
	return redisClient
}

func NewS3Client(env *configs.EnvironmentVariables) *s3.Client {
	return s3.New(s3.Options{
		Region:      env.AWSKeys.Region,
		Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(env.AWSKeys.AccessID, env.AWSKeys.SecretKey, "")),
	})
}
