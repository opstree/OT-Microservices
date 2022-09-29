VERSION := v3.0
IMG := opstree/empms-attendance:$(VERSION)

build:
	go build

build-image:
	docker build -t ${IMG} -f Dockerfile .

push-image:
	docker push ${IMG}
