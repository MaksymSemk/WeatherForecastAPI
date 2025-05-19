# Start from the official Golang image for building the app
FROM golang:1.24.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o weatherforecastapi ./cmd/main.go

EXPOSE 8080
CMD ["./weatherforecastapi"]