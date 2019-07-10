FROM golang
MAINTAINER Copyright @ airdb.com

ENV PROJECT github.com/airdb/passport
ENV WORKDIR /go/src/${PROJECT}
ENV GO111MODULE on
ADD ./ ${WORKDIR}
RUN cd ${WORKDIR} && \
	go build -o /srv/main main.go

EXPOSE 8080

WORKDIR  /srv
CMD ["/srv/main"]
