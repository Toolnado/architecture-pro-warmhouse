FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o temperature-api cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/temperature-api .

EXPOSE 8081

CMD [ "./temperature-api" ]
