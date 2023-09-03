## Overview
This is a video streaming application, similar to YouTube, built using a microservices architecture. The application allows users to upload, stream, and search videos. It's primarily built on AWS services, Golang, and the AWS CDK, and utilizes Next.js on the client-side.

# Features
Upload videos
Stream videos
Search for videos
Architecture
The application adopts a microservices architecture and uses the following services:

## Upload Service: 
Handles video uploads and stores them in an S3 bucket. Metadata is stored in DynamoDB and an event is pushed to SQS.
Video Splitting Service: Consumes video upload events from SQS, splits the videos into chunks based on time, and stores the chunks and manifests back in an S3 bucket.
Client-side: Written in Next.js and uses video.js for streaming.
Tech Stack
AWS (S3, DynamoDB, SQS, Lambda, API Gateway)
Golang
AWS CDK
Next.js
video.js
Getting Started
Prerequisites
AWS account
Golang installed
Node.js and npm installed
AWS CDK installed
Installation
Clone the repo:

bash
Copy code
git clone https://github.com/Saswat1998/Stream-X
Change to the project directory:

bash
Copy code
cd videostreaming-service
Install the dependencies:

bash
Copy code
npm install
Deploy AWS infrastructure:

bash
Copy code
cdk deploy
Usage
Video Upload
Invoke the Upload Service API through the API Gateway to upload a video.

Video Streaming
Access the application via a web browser to stream videos.
