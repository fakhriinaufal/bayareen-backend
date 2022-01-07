FROM golang:1.17-alpine as builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]