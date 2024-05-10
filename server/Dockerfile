# Use the official Go image as a builder stage
FROM golang:1.17 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o server ./cmd

# Start a new stage from scratch
FROM golang:1.17

# Set the working directory in the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=build /app/server .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./server"]