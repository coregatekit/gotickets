FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/main cmd/api/main.go

FROM alpine:3.21.3

RUN addgroup -S appgroup && adduser -S runner -G appgroup
USER runner
WORKDIR /app

COPY --from=builder /app/bin/main .

EXPOSE 8000
CMD [ "./main" ]