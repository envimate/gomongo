FROM golang

RUN apt-get update
RUN apt-get install -y libssl-dev

ADD src/gomongo /go/src/gomongo/
ADD build.sh /go
RUN /go/build.sh

ENTRYPOINT ["/bin/sh", "-c", "exec /go/bin/gomongo -mongourl $MONGO_URL"]

EXPOSE 8080
