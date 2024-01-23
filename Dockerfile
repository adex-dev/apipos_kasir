FROM golang:1.21.6-alpine
WORKDIR /app

COPY . .

RUN go build -o apipos


CMD ./apipos