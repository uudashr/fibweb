install:
	@go install ./...

prepare-install:
	@go get -v

run:
	@fibgoweb --fibgo-addr http://localhost:9090

run-fibgo:
	@docker run -d --name fibgo-server -p 9090:8080 uudashr/fibgo

stop-fibgo:
	@docker stop fibgo-server
	@docker rm -v fibgo-server

test:
	@go test

test-cover:
	@go test

check:
	@gometalinter --deadline=15s

prepare-check:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install
