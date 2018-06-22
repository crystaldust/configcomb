SHELL := /bin/bash
all:
	# go build -ldflags '-w -extldflags "-static"' -o testbin
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o k8s-api-test
	docker build -t juzhen/k8s-api-test:v1 ./
	distribute-image.sh juzhen/k8s-api-test:v1
	kubectl apply -f ./k8s-api-test.yaml

