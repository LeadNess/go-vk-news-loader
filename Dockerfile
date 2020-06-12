FROM golang:1.10

WORKDIR /go/src/news-service
COPY . .
RUN go get -t github.com/jmoiron/sqlx \
 && go get -t github.com/lib/pq \
 && go get -t github.com/go-vk-api/vk \
 && go build -ldflags "-linkmode external -extldflags -static" -a cmd/main.go

FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/news-service/main /main
COPY config /config
CMD ["/main"]