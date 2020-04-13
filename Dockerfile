# build stage
FROM golang:alpine as build-env

COPY ./ /go/src/github.com/scouball/icecast-ldap
WORKDIR /go/src/github.com/scouball/icecast-ldap

RUN apk add git
RUN go get github.com/go-ldap/ldap/v3

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/icecast-ldap


# final stage
FROM scratch
WORKDIR /go/bin
COPY --from=build-env /go/bin/ /go/bin

ENTRYPOINT ["/go/bin/icecast-ldap"]