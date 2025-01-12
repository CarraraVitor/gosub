# syntax=docker/dockerfile:1

FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /gosub
COPY go.mod go.sum ./
RUN go mod download -x 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main . 


FROM alpine:latest

RUN apk --no-cache add ca-certificates curl 
WORKDIR /gosub
COPY --from=builder /gosub/main main
COPY ./static ./static/
COPY ./subway/templates/ ./subway/templates/
COPY ./subway/style/ ./subway/style/
EXPOSE 8000

CMD ["./main"]
