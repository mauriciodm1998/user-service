run-app:
	go run cmd/user-service/main.go

run-tests-in-docker:
	docker build -t user-service . -f docker/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

