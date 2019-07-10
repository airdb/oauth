FROM golang
MAINTAINER Copyright @ airdb.com

env PROJECT github.com/airdb/passport
ENV GO111MODULE on
RUN go get ${PROJECT} && \
	cd src/${PROJECT} && \
	go build -o main main.go

# The second and final stage
FROM scratch

# Copy the binary from the builder stage
COPY --from=0 /go/src/${PROJECT}/main /srv/
COPY --from=0 /go/src/${PROJECT}/config /srv/config

EXPOSE 8080

WORKDIR  /srv
CMD ["/srv/main"]
