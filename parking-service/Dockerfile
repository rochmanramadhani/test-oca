FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./parking-service ./parking-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/parking-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/parking-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/parking-service /usr/bin/parking-service

CMD ["parking-service"]
