FROM golang:1.10-stretch as build

WORKDIR /go/src/github.com/KuaishouContainerService/quota-order-webhook
COPY . .

RUN go build -o /go/bin/quota-order-webhook cmd/*.go

FROM debian:stretch-slim

COPY --from=build /go/bin/quota-order-webhook /usr/bin/quota-order-webhook

CMD ["quota-order-webhook"]