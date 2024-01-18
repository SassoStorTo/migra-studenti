# Use Golang official image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o za-wardo

# Expose the port that the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./za-wardo"]
