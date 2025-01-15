# Choose the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the code to the container
COPY . .

# Download dependencies and compile
RUN go mod tidy
RUN go build -o main .

# Define the input command
CMD ["./main"]
