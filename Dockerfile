# Start from a Go image
FROM golang:1.22.4-alpine

# Setup Work Directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o mailgolang .

# Expose the port your app will run on
EXPOSE 3005


# Run the application
CMD ["./mailgolang"]
