FROM icejudge/golang-node

WORKDIR /go/src/github.com/sunho/fws

ADD . .

RUN make

FROM docker:stable-git

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh libc6-compat

WORKDIR /fws

COPY --from=0 /go/src/github.com/sunho/fws/out .

CMD ["/fws/server"]
