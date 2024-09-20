FROM --platform=$BUILDPLATFORM golang:1.22 as builder

## GOLANG env
ARG GOPROXY="https://proxy.golang.org,direct"
ARG GO111MODULE="on"

# Copy go.mod and download dependencies
WORKDIR /docker-test
COPY go.mod .
COPY go.sum .
COPY main.go .
RUN go mod download -x

ARG CGO_ENABLED=0
ARG TARGETOS TARGETARCH
ARG GOOS=$TARGETOS
ARG GOARCH=$TARGETARCH

COPY . .
RUN go run main.go