# syntax=docker/dockerfile:1
FROM golang:1.16
COPY . ./
RUN ls -lh
RUN GOPATH= go get \
  && GOPATH= CGO_ENABLED=0 GOOS=linux go build -a

CMD guillotine-api