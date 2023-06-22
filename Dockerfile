# Use an official Go runtime as the base image
FROM golang:1.20.4

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main cmd/main.go

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]