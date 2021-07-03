.PHONY: all build tag push rollout-restart

all: build tag push rollout-restart

build:
	docker build -t dignix/saas-honeycomb-ordermgr:latest .

tag:
	docker tag dignix/saas-honeycomb-ordermgr:latest hzhyvinskyi/saas-honeycomb-ordermgr:latest

push:
	docker push hzhyvinskyi/saas-honeycomb-ordermgr:latest

rollout-restart:
	kubectl rollout restart deployment/saas-honeycomb-ordermgr
