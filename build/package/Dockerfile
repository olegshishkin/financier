FROM golang:1.20.2-alpine3.17 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /financier ./cmd/web

FROM alpine:3.17.2
ENV	TZ=UTC
WORKDIR /
COPY --from=build /financier /financier
COPY --from=build /app/config/files /config
ARG PORTS
EXPOSE ${PORTS}
RUN adduser -D nonroot
USER nonroot
ENTRYPOINT ["/financier"]