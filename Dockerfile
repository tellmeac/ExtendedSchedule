# Stage 1: build
FROM golang:1.19.5-alpine as builder

WORKDIR /service

RUN apk update && apk add --virtual build-dependencies build-base gcc wget git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go generate ./...

RUN go test --short ./...

ENV CGO_ENABLED=0
RUN go build -a -buildvcs=false -installsuffix cgo -o app .

# Stage 2: base image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /service/app /usr/bin/app

CMD ["/usr/bin/app"]
