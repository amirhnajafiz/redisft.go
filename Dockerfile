# using alpine image with golang 1.19
FROM golang:1.19-alpine as builder

# create workdirectory
WORKDIR /src/app/

# copy go.mod and go.sum
COPY go.mod go.sum ./

# download modules
RUN go mod download

# copy all files
COPY . .

# build golang application
RUN go build -o ./main

# running stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /main .

CMD ./main