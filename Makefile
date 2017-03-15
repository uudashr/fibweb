install:
	@go install -v ./...

prepare-install:
	@go get -v ./...

run:
	@fibgoweb -fibgo-addr http://localhost:9090

docker-net-init:
	@docker network create fibnet_back-tier

docker-net-clean:
	@docker network rm fibnet_back-tier

docker-run-fibgo:
	@docker run -d --name fibgo-server --network fibnet_back-tier uudashr/fibgo

docker-stop-fibgo:
	@docker stop fibgo-server
	@docker rm -v fibgo-server

docker-build:
	@docker build -t fibweb .

setup-docker-run: docker-net-init docker-run-fibgo

teardown-docker-run: docker-stop-fibgo docker-net-clean

docker-run:
	@docker run -it --rm -p 8080:8080 --network fibnet_back-tier -e FIBGO_ADDR=fibgo-server:8080 fibweb

docker-console:
	@docker run -it --rm -p 8080:8080 --network fibnet_back-tier -e FIBGO_ADDR=fibgo-server:8080 fibweb /bin/sh

test:
	@go test ./...

test-cover:
	@go test -cover ./...

check:
	@gometalinter --deadline=15s

prepare-check:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install
