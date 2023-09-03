package server

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type VideoRequestBody struct {
	VideoName string `json:"videoName"`
}

type PreSignedURLResponse struct {
	PreSignedURL string `json:"preSignedURL"`
}

func (s *Server) UploadVideo(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body VideoRequestBody
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, err
	}

	req, _ := s.s3.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("x-stream-videos"),
		Key:    aws.String("uploads/" + body.VideoName),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to get the Pre-signed URL",
		}, err
	}
	response := PreSignedURLResponse{
		PreSignedURL: urlStr,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to Marshal reaponse",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "http://localhost:3000",
			"Access-Control-Allow-Methods": "POST, GET, PUT, OPTIONS",
		},
		Body: string(responseJSON),
	}, nil

}
