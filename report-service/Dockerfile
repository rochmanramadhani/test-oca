FROM golang:1.22.4-alpine3.20 as builder

WORKDIR /app
COPY ./lib ./lib
COPY ./report-service ./report-service

WORKDIR /app/lib
RUN go mod download

WORKDIR /app/report-service
RUN go mod download
RUN CGO_ENABLED=0 go build -o /usr/bin/report-service ./cmd

FROM alpine

COPY --from=builder /usr/bin/report-service /usr/bin/report-service

CMD ["report-service"]
