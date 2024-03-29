test-build-bake:
	docker build -t docker.io/mauricio1998/user-service . -f build/Dockerfile

run-tests:
	go test ./... -coverprofile cover.out -tags=test && go tool cover -html=cover.out

run-db:
	docker-compose -f build/db-docker-compose.yaml up -d

docker-push:
	docker push docker.io/mauricio1998/user-service

boilerplate: test-build-bake docker-push