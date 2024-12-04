FROM golang:1.23

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o main ./cmd/

CMD ["./main"]