FROM ubuntu:24.04

WORKDIR /app
COPY bin/gotk-example-http .

EXPOSE 8080

ENTRYPOINT ["/app/gotk-example-http", "http"]