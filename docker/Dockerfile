FROM golang as runner

RUN mkdir app
COPY dist/user-service app/
RUN chmod +x app
COPY ./internal/config/config.yaml /app
WORKDIR app

ENTRYPOINT [ "./user-service", "--config-dir", "." ]