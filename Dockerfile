FROM golang:alpine AS builder
COPY . .
RUN go vet .
RUN go build -o /app

FROM alpine
RUN mkdir /data
COPY --from=builder /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
