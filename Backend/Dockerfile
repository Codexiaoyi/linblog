FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

RUN go build -ldflags="-s -w" -o /bin/main main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata

COPY --from=builder /bin /app
COPY --from=builder /src/config.ini /app/config.ini
COPY --from=builder /src/controllers /app/controllers

WORKDIR /app

EXPOSE 5002

CMD ["./main"]