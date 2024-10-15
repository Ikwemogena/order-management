FROM golang:1.23-alpine AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o orders .

FROM alpine:3.18
WORKDIR /app
COPY --from=build /app/orders .

EXPOSE 8081
CMD ["./orders"]
