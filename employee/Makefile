VERSION := v3.0
IMG := opstree/empms-employee:$(VERSION)

build:
	go build

image:
	docker build -t ${IMG} -f Dockerfile .

push:
	docker push ${IMG}
