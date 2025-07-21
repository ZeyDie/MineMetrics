FROM golang:1.24-alpine AS builder

RUN apk add --no-cache bash

WORKDIR /workspace

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG VERSION=0.0.0
ENV APP_VERSION=$VERSION

RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-X 'main.Version=${APP_VERSION}' -w -s" \
    -o /workspace/bin/api \
    ./cmd/api/main.go

FROM golang:1.24-alpine

RUN apk add --no-cache bash

RUN adduser -D -h /home/container container

USER container
WORKDIR /home/container

COPY --from=builder --chown=container:container /workspace/bin/api /home/api

COPY --chown=container:container entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

CMD ["/bin/bash", "/entrypoint.sh"]