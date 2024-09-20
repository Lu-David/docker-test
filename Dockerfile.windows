ARG WINDOWS_VERSION=1809

# Build the manager binary
FROM --platform=windows/amd64 golang:1.22 as builder

## GOLANG env
ARG GO111MODULE="on" 
ARG CGO_ENABLED="0" 
ARG GOOS="windows" 
ARG GOARCH="amd64"
ARG GOPROXY="https://proxy.golang.org,direct"

# Copy go.mod and download dependencies
WORKDIR /docker-test
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download -x

COPY . .
RUN go run main.go
