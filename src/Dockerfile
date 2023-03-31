FROM golang:1.19-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
ADD gdk ./gdk
RUN go build -o godoku

FROM alpine:3.14
COPY --from=build /app/godoku /app/godoku
WORKDIR /app
EXPOSE 8080
ENV GODOKU_DIFFICULTY=44
ENV GODOKU_SIZE=9
CMD ["./godoku"]