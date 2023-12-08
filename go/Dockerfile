# Start from the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp

# Expose a port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./myapp"]