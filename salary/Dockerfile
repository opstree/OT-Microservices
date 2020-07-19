FROM golang:latest as builder
MAINTAINER OpsTree Solutions
COPY ./ /go/src/salary/
WORKDIR /go/src/salary
RUN go get -v -t -d ./... \
    && go build -o salary

FROM alpine:latest
MAINTAINER OpsTree Solutions
WORKDIR /app
RUN apk add --no-cache libc6-compat curl
COPY --from=builder /go/src/salary/salary /app/
COPY config.yaml /app/config.yaml
ENV CONFIG_FILE "/app/config.yaml"
EXPOSE 8082
ENTRYPOINT ["./salary"]
