##Builder Image
FROM golang:1.18.3-alpine as builder
ENV GO111MODULE=on
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates
COPY . /ledger
WORKDIR /ledger/cmd
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/application

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.12 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

#s Run Image
FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=builder /ledger/cmd/bin/application application
EXPOSE 8888
ENTRYPOINT ["./application"]