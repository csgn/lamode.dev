FROM golang:1.22.4-alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install

COPY . .
RUN go build -tags musl -ldflags "-extldflags -static" -v -o /usr/bin/app ./cmd/collector

COPY .env .
EXPOSE 8080

CMD ["/usr/bin/app"]
