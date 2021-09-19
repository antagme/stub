build:
	docker build -t stub .
.PHONY: build

deploy_infra: build
	docker run --rm  --name stub -itp 53:53 -p 53:53/UDP stub
.PHONY: deploy_infra

launch_request_tcp:
 	docker run -it --rm --dns=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' stub) lestienne/kdig:latest -d @172.17.0.2  www.google.es
.PHONY: launch_request_tcp

clean:
	docker rmi stub
.PHONY: clean