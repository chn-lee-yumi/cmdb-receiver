FROM alpine:latest
WORKDIR /
COPY entrypoint.sh .
COPY cmdb-receiver .
ENTRYPOINT ["/bin/sh","entrypoint.sh"]