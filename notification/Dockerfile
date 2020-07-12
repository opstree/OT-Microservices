FROM golang:latest as builder
MAINTAINER OpsTree Solutions
COPY ./ /go/src/notification/
WORKDIR /go/src/notification
RUN go get -v -t -d ./... \
    && go build -o notification

FROM alpine:latest
MAINTAINER OpsTree Solutions
WORKDIR /app
RUN apk add --no-cache libc6-compat bash
COPY --from=builder /go/src/notification/notification /bin/
COPY entrypoint.sh /app/
EXPOSE 8085
ENTRYPOINT ["./entrypoint.sh"]
