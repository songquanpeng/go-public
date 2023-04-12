FROM golang AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -ldflags "-s -w -X 'go-public/common.Version=$(cat VERSION)'" -o go-public

FROM scratch

COPY --from=builder /build/go-public /
EXPOSE 7891
EXPOSE 8080
WORKDIR /app
ENTRYPOINT ["/go-public"]
