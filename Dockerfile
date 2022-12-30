FROM alpine:latest
WORKDIR /
COPY cmdb-receiver .
ENTRYPOINT ["cmdb-receiver"]