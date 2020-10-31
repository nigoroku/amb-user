FROM golang:1.14.10-alpine3.12

WORKDIR /go/src/github.com/nigoroku/amb-user
ADD . /go/src/github.com/nigoroku/amb-user

ENV GO111MODULE=on

RUN apk add --no-cache \
    alpine-sdk \
    git \
    && go get github.com/pilu/fresh

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 8081

CMD ["fresh"]