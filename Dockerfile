FROM alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash
RUN mkdir -p /go-rest-api
WORKDIR /go-rest-api
COPY ./go-rest-server /go-rest-api/go-rest-server
RUN chmod +x ./go-rest-server
COPY ./config.toml /go-rest-api/config.toml
ADD ./db /go-rest-api/db
CMD ./go-rest-server --port=4004 --host=0.0.0.0
