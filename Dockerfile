FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . . 

RUN go build -o task-scheduler ./cmd/main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/task-scheduler .

CMD ["./task-scheduler"]

