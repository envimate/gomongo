FROM golang

ADD src/gomongo /go/src/gomongo/
ADD build.sh /
RUN /build.sh

ENTRYPOINT /go/bin/gomongo

EXPOSE 8080


