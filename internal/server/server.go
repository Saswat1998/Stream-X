package server

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Server struct {
	s3 *s3.S3
}

func New() (*Server, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})
	if err != nil {
		return nil, fmt.Errorf("error initilaizing AWS Session: %s", err)
	}
	s3Client := s3.New(sess)
	serverInstance := Server{
		s3: s3Client,
	}
	// http.HandleFunc("/upload", serverInstance.upload)
	return &serverInstance, nil
}
