# prd
FROM golang:1.15.8-alpine3.13 as build

WORKDIR src/github.com/tappn/docker-k8s-sample

ENV GO111MODULE=on

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN set -eux && \
  go build -o api main.go

FROM alpine:3.13

WORKDIR /app

COPY --from=build /go/src/github.com/tappn/docker-k8s-sample/api .

RUN set -eux && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/api

EXPOSE 8080

ENTRYPOINT ["./api"]
