FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .
COPY .env .env

RUN go mod tidy
RUN go build -o stock .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/stock .
COPY --from=build /app/.env .env  

EXPOSE 8083
CMD ["./stock"]