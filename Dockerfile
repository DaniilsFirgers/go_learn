FROM golang:1.20-alpine

WORKDIR /app

COPY . .
RUN go mod download


EXPOSE 4444:4444


CMD ["go", "run", "cmd/api/main.go"]