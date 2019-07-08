FROM golang
MAINTAINER Copyright @ airdb.com

ENV GO111MODULE on
RUN go get github.com/airdb/passport && \
	go build -o main github.com/airdb/passport/main.go

# The second and final stage
FROM scratch

# Copy the binary from the builder stage
#COPY --from=0 /go/src/github.com/bbhj/minabbs/main /srv/
#COPY --from=0 /go/src/github.com/bbhj/minabbs/keys /srv/keys

EXPOSE 8080

WORKDIR  /srv
CMD ["/srv/main"]
