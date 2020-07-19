FROM golang:latest as builder
MAINTAINER OpsTree Solutions
COPY . /go/src/attendance
WORKDIR /go/src/attendance
RUN go get -v -t -d ./... \
    && go build -o attendance

FROM alpine:latest
MAINTAINER OpsTree Solutions
WORKDIR /app
RUN apk add --no-cache libc6-compat curl
COPY --from=builder /go/src/attendance/attendance /app/
COPY config.yaml /app/config.yaml
ENV CONFIG_FILE "/app/config.yaml"
EXPOSE 8081
ENTRYPOINT ["./attendance"]
