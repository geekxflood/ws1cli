# Use a specific version of golang for reproducible builds
FROM golang:1.21-alpine AS builder

# Setting environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory outside of GOPATH to support Go Modules.
WORKDIR /workspace

# Copy go.mod and go.sum to download dependencies efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary in a way that includes version information.
# The version is taken from the VERSION file in the context.
RUN go build -ldflags="-s -w -X 'geekxflood/ws1cli/internal/version.Version=$(cat VERSION)'" -o /bin/ws1cli

# Second stage: Use the official Alpine image for a lean production container.
FROM alpine:3.18

# Add CA certificates
RUN apk add --no-cache ca-certificates

# Copy the pre-built binary file from the previous stage
COPY --from=builder /bin/ws1cli /usr/local/bin/ws1cli

# Set the command to run when starting the container
ENTRYPOINT ["/usr/local/bin/ws1cli"]
