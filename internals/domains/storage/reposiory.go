package storage

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

type Repository interface {
	DeclareStorageUnit(bucketName string) error
	CreateStorageUnit(bucketName string) error
	UploadFile(bucketName string, objectKey string, file *os.File) (string, error)
}

type S3ClientInterface interface {
	HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error)
	PutBucketPolicy(ctx context.Context, params *s3.PutBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.PutBucketPolicyOutput, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}
