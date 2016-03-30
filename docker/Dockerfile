FROM golang

ADD src/gomongo /go/src/gomongo/
ADD build.sh /go
RUN /go/build.sh

ENTRYPOINT /go/bin/gomongo

EXPOSE 8080


