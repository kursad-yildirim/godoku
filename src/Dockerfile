# Start with the official Go image
FROM golang:1.17.1-alpine3.14 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go mod and sum files and download dependencies
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

# Copy the source code into the container
COPY *.go ./
ADD gdk ./gdk
# Build the binary
RUN go build -o godoku

# Use a smaller base image for the final image
FROM alpine:3.14

# Copy the binary from the build container
COPY --from=build /app/godoku /app/godoku

# Set the working directory inside the container
WORKDIR /app
EXPOSE 8080
ENV GODOKU_DIFFICULTY=44
ENV GODOKU_SIZE=9

# Run the binary
CMD ["./godoku"]