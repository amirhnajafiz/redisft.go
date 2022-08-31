# event-man dockerfile
FROM golang:1.17 as builder

WORKDIR /src/app/

COPY . .

RUN go build -o ./execute

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /execute .

ENTRYPOINT ["./execute"]
