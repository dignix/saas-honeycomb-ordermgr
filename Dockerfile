FROM golang:1.16.5-alpine3.14 AS builder
ENV GO111MODULE=on
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM scratch
COPY --from=builder /src/app /app
EXPOSE 8091
ENTRYPOINT ["/app"]
