VERSION := v3.0
IMG := opstree/empms-salary:$(VERSION)
build:
	go build

image:
	docker build -t ${IMG} -f Dockerfile .

push:
	docker push ${IMG}
