FROM alpine:latest
WORKDIR /
COPY entrypoint.sh .
COPY cmdb-receiver .
ENTRYPOINT ["entrypoint.sh"]