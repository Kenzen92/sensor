# Use the official Golang image as a base
FROM golang:1.23-alpine


# Install SQLite and build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Enable CGO
ENV CGO_ENABLED=1


# Set the working directory inside the container
WORKDIR /app

# Copy Go module files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the Go source code to the working directory
COPY . .

# Install the SQLite library
RUN apk add --no-cache sqlite

# Build the Go application
RUN go build -o server .

# Expose port 8080 for the HTTP server
EXPOSE 8080

# Run the Go HTTP server
CMD ["./server"]
