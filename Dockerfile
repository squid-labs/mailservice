FROM golang:alpine AS build

ENV GO111MODULE=on

WORKDIR /go/src/app

LABEL maintainer="github@nomadslayer"

RUN apk add bash ca-certificates git gcc g++ libc-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /go/bin/mailservice ./cmd/mailservice

FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /go/bin/mailservice ./mailservice
COPY config/ config/
COPY templates/ templates/

ENTRYPOINT ["./mailservice", "-config", "config/mailservice.yml"]
