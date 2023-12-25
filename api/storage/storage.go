package storage

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	appConfig "main/config"
	"time"
)

type Storage struct {
	client     *s3.Client
	bucketName string
}

func New(cfg appConfig.Config) (*Storage, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           cfg.S3.Endpoint,
			SigningRegion: cfg.S3.AwsDefaultRegion,
		}, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(cfg.S3.AwsDefaultRegion),
		config.WithEndpointResolverWithOptions(customResolver),
	)

	if err != nil {
		return nil, err
	}

	st := &Storage{}
	st.bucketName = cfg.S3.Bucket

	st.client = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	err = st.createBucketIfNotExists(cfg)
	if err != nil {
		return nil, err
	}

	return st, nil
}

func (receiver Storage) createBucketIfNotExists(cfg appConfig.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := receiver.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(cfg.S3.Bucket),
	})

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err != nil {
		log.Println("creating bucket")

		_, err := receiver.client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(cfg.S3.Bucket),
		})
		return err
	} else {
		log.Println("bucket already exists")
	}

	return nil
}

func (receiver Storage) Upload(key string, fileBytes []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := receiver.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(receiver.bucketName),
		Key:    aws.String(fmt.Sprint(key)),
		Body:   bytes.NewReader(fileBytes),
	})

	return err
}
