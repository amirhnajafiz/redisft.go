# event-man dockerfile
FROM golang:1.17 as builder

WORKDIR /src/app/

COPY . .

RUN make build

FROM alpine:latest

WORKDIR /app/

COPY --from=builder ./execute .

ENTRYPOINT ["./execute"]
