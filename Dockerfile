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

COPY . /app
WORKDIR /app

RUN go mod vendor
RUN go mod download
RUN go mod verify

RUN go build -tags musl -ldflags="-w -s" -o cryptobot cmd/*.go

FROM alpine:latest AS production

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /app/cryptobot /app/cryptobot
COPY --from=builder /app/config/config.yaml /app/config/config.yaml
COPY --from=builder /app/migrations/postgres /app/migrations/postgres

USER appuser:appuser

CMD ["./app/cryptobot"]