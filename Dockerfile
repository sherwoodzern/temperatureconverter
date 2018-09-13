FROM golang:1.11 as builder

RUN mkdir -p /go/src/github.com/sherwoodzern/temperatureconverter
WORKDIR  /go/src/github.com/sherwoodzern/temperatureconverter

COPY . .

RUN CGO_ENABLED=0 go build -i -installsuffix cgo -o temperatureconverter .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir -p public

COPY --from=builder /go/src/github.com/sherwoodzern/temperatureconverter .
EXPOSE 3000
ENTRYPOINT [ "./temperatureconverter" ]

