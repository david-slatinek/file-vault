package storage

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/alexedwards/argon2id"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	appConfig "main/config"
	"runtime"
	"time"
)

type Storage struct {
	client      *s3.Client
	bucketName  string
	argonParams *argon2id.Params
}

func New(cfg *appConfig.Config) (*Storage, error) {
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
	st.argonParams = &argon2id.Params{
		Memory:      128 * 1024, // 128 MB
		Iterations:  24,
		Parallelism: uint8(runtime.NumCPU()),
		SaltLength:  16,
		KeyLength:   32,
	}

	st.client = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	err = st.createBucketIfNotExists(cfg)
	if err != nil {
		return nil, err
	}

	return st, nil
}

func (receiver Storage) createBucketIfNotExists(cfg *appConfig.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := receiver.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(cfg.S3.Bucket),
	})

	if err != nil {
		log.Println("creating bucket")

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := receiver.client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(cfg.S3.Bucket),
		})
		return err
	} else {
		log.Println("bucket already exists")
	}

	return nil
}

func (receiver Storage) Upload(fileKey, password string, fileBytes []byte) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ciphertext, salt, err := receiver.encryptFile(password, fileBytes)
	if err != nil {
		return "", err
	}

	_, err = receiver.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(receiver.bucketName),
		Key:    aws.String(fmt.Sprint(fileKey)),
		Body:   bytes.NewReader(ciphertext),
	})

	return salt, err
}

func (receiver Storage) encryptFile(password string, plaintext []byte) ([]byte, string, error) {
	hash, err := argon2id.CreateHash(password, receiver.argonParams)
	if err != nil {
		return nil, "", err
	}

	_, salt, key, err := argon2id.DecodeHash(hash)
	if err != nil {
		return nil, "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, hex.EncodeToString(salt), nil
}

func (receiver Storage) Delete(fileKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := receiver.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(receiver.bucketName),
		Key:    aws.String(fileKey),
	})

	return err
}
