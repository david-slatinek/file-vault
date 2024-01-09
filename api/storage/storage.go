package storage

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/tvdburgt/go-argon2"
	"io"
	"log"
	appConfig "main/config"
	"runtime"
	"time"
)

type Storage struct {
	client      *s3.Client
	bucketName  string
	argonParams *argon2.Context
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
	st.argonParams = &argon2.Context{
		Iterations:  24,
		Memory:      128 * 1024, // 128 MB
		Parallelism: runtime.NumCPU(),
		HashLen:     32,
		Mode:        argon2.ModeArgon2id,
		Version:     argon2.Version13,
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

func randomBytes(n int) ([]byte, error) {
	key := make([]byte, n)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (receiver Storage) encryptFile(password string, plaintext []byte) ([]byte, string, error) {
	salt, err := randomBytes(16)
	if err != nil {
		return nil, "", err
	}

	hash, err := argon2.Hash(receiver.argonParams, []byte(password), salt)
	if err != nil {
		return nil, "", err
	}

	block, err := aes.NewCipher(hash)
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

func (receiver Storage) Download(fileKey, password, salt string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := receiver.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(receiver.bucketName),
		Key:    aws.String(fileKey),
	})
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(result.Body)
	if err != nil {
		return nil, err
	}

	ciphertext := buf.Bytes()

	plaintext, err := receiver.decryptFile(password, salt, ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (receiver Storage) decryptFile(password, salt string, ciphertext []byte) ([]byte, error) {
	saltBytes, err := hex.DecodeString(salt)
	if err != nil {
		return nil, err
	}

	hash, err := argon2.Hash(receiver.argonParams, []byte(password), saltBytes)

	block, err := aes.NewCipher(hash)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
