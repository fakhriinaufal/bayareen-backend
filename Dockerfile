FROM golang:1.17-alpine as builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main

FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/main .
RUN mkdir config
COPY --from=builder /app/config/config.toml ./config
RUN mkdir -p features/transaction/service/template
COPY --from=builder /app/features/transaction/service/template/*.html ./features/transaction/service/template
CMD ["./main"]
