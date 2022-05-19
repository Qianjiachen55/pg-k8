## BUild stage

FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY db/migration ./migration
COPY tools ./tools
COPY wait-for .


EXPOSE 8000
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]
