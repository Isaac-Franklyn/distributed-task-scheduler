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

# At the bottom:
COPY wait-for-cockroach.sh .

RUN chmod +x wait-for-cockroach.sh

# Use entrypoint to wait for the port to open
ENTRYPOINT ["./wait-for-cockroach.sh", "cockroach", "26257", "./task-scheduler"]


