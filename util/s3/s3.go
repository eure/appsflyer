package s3

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Uploader s3 uploader
type Uploader struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	Bucket          string
}

// Do upload file to s3
func (u Uploader) Do(file *os.File) error {
	cli := s3.New(session.New(), &aws.Config{
		Credentials: credentials.NewStaticCredentials(
			u.AccessKeyID,
			u.SecretAccessKey,
			""),
		Region: aws.String(u.Region),
	})
	_, fileName := filepath.Split(file.Name())
	if _, err := cli.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.Bucket),
		Key:    aws.String(fileName),
		Body:   file,
	}); err != nil {
		return err
	}
	return nil
}
