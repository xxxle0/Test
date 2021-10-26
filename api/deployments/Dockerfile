FROM golang:1.15-alpine AS builder

RUN set -eux && \
    apk add --no-progress --no-cache \
        gcc \
        musl-dev

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

COPY . .

RUN go mod vendor
RUN go mod download

RUN go build -a -tags musl -o /src/bin/main .

FROM alpine:3.9 

WORKDIR /app
EXPOSE 8080

RUN apk add ca-certificates

COPY --from=builder /src/bin/main /app/
RUN mkdir /app/resources

ENTRYPOINT ["/app/main"]
