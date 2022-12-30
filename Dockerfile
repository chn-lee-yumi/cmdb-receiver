FROM golang:1.19-alpine AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o cmdb-receiver .

FROM alpine:latest
WORKDIR /
COPY --from=build /app/cmdb-receiver /cmdb-receiver
ENTRYPOINT ["/cmdb-receiver"]