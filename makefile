run:
	go run main.go

build-local:
	docker build -f build/development/Dockerfile  -t new-proto-dev:latest .
	docker run -p 8888:8888 new-proto-dev

mockgen:
	cd api/services && mockery --all
	cd api/repository && mockery --all

test:
	go test -coverpkg ./... ./... | { grep -v 'no test files'; true; }

coverage:
	go test -coverprofile=coverage.out -coverpkg  ./... ./... | { grep -v 'no test files'; true; }
	go tool cover -html=coverage.out

