FROM golang:1.23.0-alpine3.19 as modules

COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.23.0-alpine3.19 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./cmd/app

FROM scratch
COPY --from=builder /app/config /config
COPY --from=builder /bin/app /app
EXPOSE 8080
CMD ["/app"]