FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
COPY .env .env

RUN go mod tidy
RUN go build -o gateway .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/gateway .
COPY --from=build /app/.env .env 

EXPOSE 8080
CMD ["./gateway"]