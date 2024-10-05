FROM golang:1.22.4-alpine AS builder
WORKDIR /app

RUN apk -U add ca-certificates
RUN apk update && apk upgrade && apk add pkgconf git bash build-base sudo
RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 GOOS=linux GOARACH=amd64 GO111MODULE=on go build -tags musl -ldflags "-extldflags -static" -v -o /go/bin/collector .

FROM scratch
WORKDIR /app
COPY --from=builder /go/bin/collector /app/collector

EXPOSE ${APP_PORT}

ENTRYPOINT ["/app/collector"]
