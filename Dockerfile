# Use an official Go runtime as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY go.mod .
COPY server.go .

# Build the Go binary
RUN go get
RUN go build -o bin .

# Expose the port your application listens on
Entrypoint [`/app/bin`]

# Command to run your application
CMD ["./myapp"]
