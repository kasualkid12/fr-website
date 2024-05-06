package encryption

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
)

var KMSClient *kms.Client

func init() {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatal("AWS region not set in the environment variables")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	KMSClient = kms.NewFromConfig(cfg)
}

func EncryptWithKMS(data string, keyAlias string) (string, error) {
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

func DecryptWithKMS(encryptedData string, keyAlias string) (string, error) {
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