FROM iron/go:dev

MAINTAINER Opstree Solutions

WORKDIR /app

ENV SRC_DIR=/go/src/gitlab.com/opstree/ot-go-webapp/

ADD . $SRC_DIR

RUN cd $SRC_DIR; go get -v -t -d ./... && \
    go build -o myapp; cp myapp /app/

ENTRYPOINT ["./myapp"]
