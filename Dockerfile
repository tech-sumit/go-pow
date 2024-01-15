# Use an official Golang runtime as a parent image
FROM golang:1.21

# Set the working directory in the container
WORKDIR /app

# Copy the server and client code into the container
COPY . .

# Build the server and client
RUN go build -o server cmd/server/server.go
RUN go build -o client cmd/client/client.go

# Expose the necessary ports
EXPOSE 8080

# Run the server and client as separate processes
CMD ./server & sleep 5 && ./client
