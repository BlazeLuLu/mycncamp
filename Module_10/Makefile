export tag=v1.0

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t blazelulu/httpserver:${tag} .

push: release
	echo "pushing blazelulu/httpserver"
	docker push blazelulu/httpserver:v1.0