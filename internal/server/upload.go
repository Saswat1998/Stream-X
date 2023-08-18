package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type VideoRequestBody struct {
	VideoData string `json:"videoData"`
	VideoName string `json:"videoName"`
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

	videoData := body.VideoData
	videoBytes, err := base64.StdEncoding.DecodeString(videoData)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error decoding base64 video data",
		}, err
	}
	_, err = s.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("x-stream-videos"),
		Key:    aws.String("uploads/" + body.VideoName),
		Body:   bytes.NewReader(videoBytes),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error uploading the video",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Video uploaded successfully!",
	}, nil

}
