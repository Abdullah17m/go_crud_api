# Use Go official image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install Git and dependencies
RUN apk add --no-cache git

# Copy go mod and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the code
COPY . .

# Build the app
RUN go build -o server main.go

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./server"]
