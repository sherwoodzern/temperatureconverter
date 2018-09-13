FROM golang:1.9.2

RUN mkdir /conversions
COPY . /go/src/temperatureconvert

RUN go get -u github.com/FiloSottile/gvt
RUN cd /go/src/temperatureconvert && gvt restore

RUN CGO_ENABLED=0 GOOS=linux go build -o /conversions/main temperatureconvert/cmd

CMD ["/conversions/main"]

