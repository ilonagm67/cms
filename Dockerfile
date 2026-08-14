FROM golang:alpine AS builder
COPY . .
RUN go vet .
RUN go build -o /app

FROM alpine
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
