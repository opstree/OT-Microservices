FROM golang:latest as builder
MAINTAINER OpsTree Solutions
COPY ./ /go/src/employee
WORKDIR /go/src/employee
RUN go get -v -t -d ./... \
    && go build -o employee

FROM alpine:latest
MAINTAINER OpsTree Solutions
WORKDIR /app
RUN apk add --no-cache libc6-compat curl
COPY --from=builder /go/src/employee/employee /app/
COPY ./config.yaml /app/config.yaml
ENV CONFIG_FILE "/app/config.yaml"
EXPOSE 8083
ENTRYPOINT ["./employee"]
