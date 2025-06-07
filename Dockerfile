FROM golang:alpine AS builder
WORKDIR /bot
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bot/main /main
WORKDIR /
ENTRYPOINT ["/main"]