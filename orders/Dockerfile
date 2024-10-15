FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
COPY .env .env

RUN go mod tidy
RUN go build -o orders .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/orders .
COPY --from=build /app/.env .env 

EXPOSE 8081
CMD ["./orders"]
