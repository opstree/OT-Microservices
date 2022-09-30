VERSION := v3.0
IMG := opstree/empms-gateway:$(VERSION)

build:
	mvn package

image:
	docker build -t ${IMG} -f Dockerfile .

push:
	docker push ${IMG}
