##Builder Image
FROM golang:1.18.3-stretch as builder
ENV GO111MODULE=on
COPY . /ledger
WORKDIR /ledger
RUN go mod download
RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -o bin/application

#s Run Image
FROM scratch
COPY --from=builder /ledger/bin/application application
EXPOSE 8080
ENTRYPOINT ["./application"]