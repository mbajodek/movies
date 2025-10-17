FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apk update && apk add make
RUN make build

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /app/bin/movies/main .

EXPOSE 8080

CMD ["./main"]