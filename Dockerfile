FROM golang:1.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest AS production
COPY --from=builder /app .
  
ENV DB_USERNAME=postgres
ENV DB_PASSWORD=postgres
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_TABLE=postgres
ENV JWT_SECRET=secret
ENV REFRESH_TOKEN_SECRET=secret

CMD ["./app"]
