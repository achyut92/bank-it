# Use official Golang image
FROM golang:latest


# Set working directory
WORKDIR /app

# Copy go.mod and go.sum, download dependencies first (layer cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire app
COPY . .

# Build the app
RUN go build -o app .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./app"]
