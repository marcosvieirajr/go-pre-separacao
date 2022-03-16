# syntax=docker/dockerfile:1

FROM alpine:3.15 as authority
RUN mkdir /user && \
    echo 'appuser:x:1000:1000:appuser:/:' > /user/passwd && \
    echo 'appgroup:x:1000:' > /user/group
RUN apk --no-cache add ca-certificates


FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags nomsgpack -ldflags="-s -w" -a -o /server ./cmd/web


FROM scratch
LABEL maintainer="marcosvieirajr@gmail.com"
WORKDIR /
COPY --from=authority /user/group /user/passwd /etc/
COPY --from=authority /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /server .
ENV HOST_PORT=8080
EXPOSE 8080
USER appuser:appgroup
ENTRYPOINT ["/server"]
