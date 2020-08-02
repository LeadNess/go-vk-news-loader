FROM golang:1.10 as builder

WORKDIR /go/src/news-service
COPY . .
RUN go get -t github.com/LeadNess/go-vk-news-loader \
 && go build -ldflags "-linkmode external -extldflags -static" -a cmd/main.go

FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM scratch
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/news-service/main /main
COPY config /config
CMD ["/main"]