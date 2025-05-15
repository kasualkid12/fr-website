package encryption

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

var KMSClient *kms.Client
var keyAlias string

func init() {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
	if region == "" {
		log.Fatal("AWS region not set in the environment variables")
	}

	keyAlias = os.Getenv("KEY_ALIAS")
	if keyAlias == "" {
		log.Fatal("Key alias not set in the environment variables")
	}

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("KMS_ENDPOINT")

	var cfg aws.Config
	var err error

	if endpoint != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           endpoint,
						SigningRegion: region,
					}, nil
				},
			)),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
		)
	}
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	KMSClient = kms.NewFromConfig(cfg)
}

func EncryptWithKMS(data string) (string, error) {
	input := &kms.EncryptInput{
		KeyId:     aws.String(keyAlias),
		Plaintext: []byte(data),
	}

	result, err := KMSClient.Encrypt(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt data: %v", err)
	}

	return base64.StdEncoding.EncodeToString(result.CiphertextBlob), nil
}

func DecryptWithKMS(encryptedData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", fmt.Errorf("failed to decode data: %v", err)
	}

	input := &kms.DecryptInput{
		KeyId:          aws.String(keyAlias),
		CiphertextBlob: data,
	}

	result, err := KMSClient.Decrypt(context.Background(), input)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt data: %v", err)
	}

	return string(result.Plaintext), nil
}
