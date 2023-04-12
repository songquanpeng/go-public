FROM golang AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -X 'go-public/common.Version=$(cat VERSION)'" -o go-public

FROM alpine

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates 2>/dev/null || true

COPY --from=builder /build/go-public /
EXPOSE 6871
EXPOSE 8080
WORKDIR /app
ENTRYPOINT ["/go-public"]

