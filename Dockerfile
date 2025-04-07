FROM golang:1.21-alpine

# Set up workspace
WORKDIR /app
COPY . .

# Build the Go app
RUN go build -o app .

# Run the app
CMD ["./app"]