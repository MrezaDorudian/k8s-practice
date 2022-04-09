## first stage

FROM golang:alpine3.15 as builder

LABEL maintainer="mr.dorudian@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN go build -o main

EXPOSE 8080

## second stage

FROM alpine:3.15

RUN apk add --update --no-cache ca-certificates

WORKDIR /app/

COPY --from=builder /app/main .

CMD ["./main"]