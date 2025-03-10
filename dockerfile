# Use the official Golang image as a build stage
FROM golang:1.24


WORKDIR /app

COPY go.mod .
COPY go.sum .
#COPY main.go .
COPY . .

RUN go get

RUN go build -o bin .


ENTRYPOINT ["/app/bin"]