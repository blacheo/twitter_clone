FROM golang:1.21

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

EXPOSE 5000



CMD go run *.go

