# Step 1: Build
FROM golang:1.15-alpine as build
WORKDIR /go/app
COPY . .

RUN go build -o latest

# Step 2: Package for Run
FROM alpine
WORKDIR /app

COPY --from=build /go/app/latest .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/latest

CMD ["./latest"]
