FROM ubuntu:14.04

RUN apt-get update -y && apt-get install -y curl git golang

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

WORKDIR /go/src/github.com/mattmanning/geoip

RUN curl -O http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz
RUN gunzip GeoLite2-City.mmdb.gz

COPY . /go/src/github.com/mattmanning/geoip

RUN go get github.com/gorilla/mux
RUN go get github.com/oschwald/geoip2-golang
RUN go install ./...

CMD ["geoip"]
