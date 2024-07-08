FROM golang:1.21.5-alpine3.17 as builder

RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo

ENV USER=appuser
ENV UID=10001
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" "${USER}"

COPY . /crypto-collector
WORKDIR /crypto-collector

RUN go mod vendor
RUN go mod download
RUN go mod verify

RUN go build -tags musl -ldflags="-w -s" -o crypto-collector cmd/*.go

FROM alpine:latest AS production

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /crypto-collector/crypto-collector /crypto-collector/crypto-collector
COPY --from=builder /crypto-collector/config/config.yaml /crypto-collector/config/config.yaml
COPY --from=builder /crypto-collector/migrations/postgres /crypto-collector/migrations/postgres

USER appuser:appuser

CMD ["./crypto-collector/crypto-collector"]