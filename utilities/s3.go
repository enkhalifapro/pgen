package utilities

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/viper"
)

// S3Adapter describes s3 client.
type S3Adapter interface {
	// Put content from body to key, last argument can be nil.
	Put(ctx context.Context, key string, body io.ReadSeeker, metadata map[string]*string) error

	// Get object by key.
	Get(ctx context.Context, key string) (io.ReadCloser, map[string]*string, error)
}

// S3 client wraper realizes S3Adapter interface.
type S3 struct {
	session *session.Session
}

// getSession for connection. This method cache connection if it already was created.
func (s *S3) getSession() (*session.Session, error) {
	if s.session == nil {
		cfg := aws.Config{
			Credentials: credentials.NewStaticCredentials(
				viper.GetString("s3.id"),
				viper.GetString("s3.secret"),
				"",
			),
			Region: aws.String(viper.GetString("s3.region")),
		}

		session, err := session.NewSessionWithOptions(session.Options{Config: cfg})
		if err != nil {
			return nil, err
		}
		s.session = session
	}

	return s.session, nil
}

// Put content from body to key, last argument can be nil.
func (s *S3) Put(ctx context.Context, key string, body io.ReadSeeker, metadata map[string]*string) error {
	sess, err := s.getSession()
	if err != nil {
		return err
	}

	_, err = s3.New(sess).PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket:   aws.String(viper.GetString("s3.bucket.history")),
		Key:      aws.String(key),
		Body:     body,
		Metadata: metadata,
		ACL:      aws.String("public-read"),
	})
	return err
}

// Get object by key.
func (s *S3) Get(ctx context.Context, key string) (io.ReadCloser, map[string]*string, error) {
	sess, err := s.getSession()
	if err != nil {
		return nil, nil, err
	}

	obj, err := s3.New(sess).GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(viper.GetString("s3.bucket.history")),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, nil, err
	}

	return obj.Body, obj.Metadata, nil
}
