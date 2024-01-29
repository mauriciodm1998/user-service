FROM golang as builder

RUN mkdir app
COPY ./ app
WORKDIR app
RUN CGO_ENABLED=0 go test ./... -coverprofile cover.out -tags=test && go build -o dist/user-service cmd/user-service/main.go

FROM golang as runner

RUN mkdir app
COPY --from=builder ./go/app/dist/user-service app/
RUN chmod +x app
WORKDIR app

ENTRYPOINT [ "./user-service", "--config-dir", "." ]