FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

RUN mkdir -p /app
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app

FROM gcr.io/distroless/static
 
COPY --from=builder /app /app

EXPOSE 8080
ENTRYPOINT ["/app/darwin-proxy"]