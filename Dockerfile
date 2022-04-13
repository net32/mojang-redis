FROM golang:1.18-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Isaias <contato@isaias.dev>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]