FROM alpine:latest AS build
RUN apk update
RUN apk upgrade
RUN apk add --update go gcc g++
WORKDIR /app
ENV GOPATH /app
ADD . /app/src/cmdb-receiver
RUN CGO_ENABLED=1 GOOS=linux go build cmdb-receiver

FROM alpine:latest
WORKDIR /
COPY --from=build /app/bin/cmdb-receiver /cmdb-receiver
ENTRYPOINT ["/cmdb-receiver"]