FROM golang:1.15-alpine as build

WORKDIR /go/app
COPY . .

RUN go build -o app

FROM alpine

WORKDIR /app

COPY --from=build /go/app/app .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/app

CMD ["./app"]
