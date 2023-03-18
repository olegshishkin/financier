FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /financier

FROM alpine:3.17.2
COPY --from=build /financier /financier
COPY --from=build /app/config /config
EXPOSE 8080
RUN adduser -D nonroot
USER nonroot
ENTRYPOINT ["/financier"]