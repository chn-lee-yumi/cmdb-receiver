FROM alpine:latest
WORKDIR /
COPY entrypoint.sh /entrypoint.sh
COPY cmdb-receiver /cmdb-receiver
ENTRYPOINT ["/bin/sh","entrypoint.sh"]