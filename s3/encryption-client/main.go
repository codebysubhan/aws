package main

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os" // Added to read local files

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// SimplifiedEncryptionClient mimics the Ruby Aws::S3::Encryption::Client architecture
type SimplifiedEncryptionClient struct {
	S3Client *s3.Client
	AESKey   []byte
}

// PutObject encrypts locally on your system and drops it in S3
func (sec *SimplifiedEncryptionClient) PutObject(ctx context.Context, bucket, key, body string) (*s3.PutObjectOutput, error) {
	block, err := aes.NewCipher(sec.AESKey)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Local client side encryption step
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(body), nil)

	return sec.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(ciphertext),
	})
}

// GetObject grabs from S3 and decrypts locally on your system
func (sec *SimplifiedEncryptionClient) GetObject(ctx context.Context, bucket, key string) (string, error) {
	result, err := sec.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	ciphertext, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(sec.AESKey)
	if err != nil {
		return "", err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, actualCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, actualCiphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func main() {
	ctx := context.Background()
	bucket := "encrypt-fun-sa-12321"
	objectKey := "secret"

	// 1. Read from your local hello.txt file
	log.Println("Reading local file contents from hello.txt...")
	fileContent, err := os.ReadFile("hello.txt")
	if err != nil {
		log.Fatalf("failed to read hello.txt file: %v", err)
	}

	// 2. Setup standard AWS S3 client configurations
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load SDK config: %v", err)
	}
	standardS3 := s3.NewFromConfig(cfg)

	// 3. Define your static local master key (32 bytes for AES-256)
	localKey := []byte("a-very-secret-32-byte-long-key!!")

	// 4. Initialize your simplified encryption wrapper
	s3EncryptionClient := &SimplifiedEncryptionClient{
		S3Client: standardS3,
		AESKey:   localKey,
	}

	// ==========================================
	// ROUND-TRIP: PUT AND GET WITH THE ENCRYPTION CLIENT
	// ==========================================
	fmt.Println("\n--- PUT via Encryption Client ---")
	// Passing fileContent instead of the hardcoded string
	_, err = s3EncryptionClient.PutObject(ctx, bucket, objectKey, string(fileContent))
	if err != nil {
		log.Fatalf("PUT error: %v", err)
	}
	fmt.Println("Status: Success (hello.txt contents securely uploaded!)")

	fmt.Println("\n--- GET WITH KEY via Encryption Client ---")
	decryptedBody, err := s3EncryptionClient.GetObject(ctx, bucket, objectKey)
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}
	fmt.Printf("Result: %s", decryptedBody) // Should log => 'Hello marsss'

	// ==========================================
	// GET WITHOUT KEY (Directly using standard SDK client)
	// ==========================================
	fmt.Println("\n--- GET WITHOUT KEY via Standard S3 Client ---")
	rawResult, err := standardS3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Fatalf("Standard GET error: %v", err)
	}
	defer rawResult.Body.Close()

	rawBytes, _ := io.ReadAll(rawResult.Body)
	fmt.Printf("Resulting Raw Cipher Text: %v\n", rawBytes)
}