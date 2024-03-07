FROM golang:latest

RUN go install github.com/cosmtrek/air@latest

WORKDIR /var/www/

RUN go env -w GOFLAGS="-buildvcs=false"

ENTRYPOINT [ "air" ]