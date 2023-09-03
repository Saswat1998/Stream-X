# Video Streaming Service

## Overview

This is a video streaming application, similar to YouTube, built using a microservices architecture. The application allows users to upload, stream, and search videos. It's primarily built on AWS services, Golang, and the AWS CDK, and utilizes Next.js on the client-side.

## Table of Contents

- [Features](#features)
- [Architecture](#architecture)
  - [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)

## Features

- **Upload Videos**: Users can upload videos.
- **Stream Videos**: Users can stream videos.
- **Search Videos**: Users can search for specific videos.

## Architecture

The application adopts a microservices architecture and uses the following services:

### Upload Service
- Handles video uploads and stores them in an S3 bucket.
- Metadata is stored in DynamoDB.
- Pushes an event to SQS after a successful upload.

### Video Splitting Service
- Consumes video upload events from SQS.
- Splits the videos into chunks based on time.
- Stores the chunks and manifests back in an S3 bucket.

### Client-side
- Written in Next.js.
- Uses video.js for video streaming capabilities.

### Tech Stack
- AWS (S3, DynamoDB, SQS, Lambda, API Gateway)
- Golang
- AWS CDK
- Next.js
- video.js

## Getting Started

### Prerequisites
- AWS account
- Golang installed
- Node.js and npm installed
- AWS CDK installed

### Installation

1. **Clone the repo**:
    ```bash
    git clone https://github.com/Saswat1998/Stream-X.git
    ```

2. **Navigate to project directory**:
    ```bash
    cd videostreaming-service
    ```

3. **Install Dependencies**:
    ```bash
    npm install
    ```

4. **Deploy AWS Infrastructure**:
    ```bash
    cdk deploy
    ```

### Usage

#### Video Upload
Invoke the Upload Service API through the API Gateway to upload a video.

#### Video Streaming
Access the application via a web browser to stream videos.


Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
