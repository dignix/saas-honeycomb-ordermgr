APP_NAME := saas-honeycomb-ordermgr
IMAGE_TAG := latest

.PHONY: all build tag push rollout-restart

all: build tag push rollout-restart

build:
	docker build --rm -t dignix/$(APP_NAME):$(IMAGE_TAG) .

tag:
	docker tag dignix/$(APP_NAME):$(IMAGE_TAG) hzhyvinskyi/$(APP_NAME):$(IMAGE_TAG)

push:
	docker push hzhyvinskyi/$(APP_NAME):$(IMAGE_TAG)

rollout-restart:
	kubectl rollout restart deployment/$(APP_NAME) -n ordermgr
