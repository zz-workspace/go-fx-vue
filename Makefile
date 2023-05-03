init:
	./init.sh
build:
	ENV=production docker compose build client --no-cache
start:
	alias air='$(go env GOPATH)/bin/air'
	air
clean:
	rm -rf golang-my-app
	docker service rm service-08
	