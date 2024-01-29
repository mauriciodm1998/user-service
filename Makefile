test-build-bake:
	docker build -t user-service . -f build/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

run-db:
	docker-compose -f build/db-docker-compose.yml up -d

run-app:
	go run cmd/user-service/main.go