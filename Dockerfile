# Use the official Go image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o parser

# Expose the port on which the application will run (if applicable)
ENV PORT=80
ENV INITIAL_BLOCK=19476043

# Define the command to run the application
CMD ["./parser"]