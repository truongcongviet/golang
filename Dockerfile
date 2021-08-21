FROM golang:alpine as builder


WORKDIR /friend-management-v1

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build

CMD ["./friend-management-v1"]