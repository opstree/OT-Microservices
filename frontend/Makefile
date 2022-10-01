VERSION := v3.0
IMG := opstree/empms-frontend:$(VERSION)

install:
	npm install

build: install
	npm run build

image:
	docker build -t ${IMG} -f Dockerfile .

push:
	docker push ${IMG}
