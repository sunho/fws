FROM icejudge/golang-node

WORKDIR /go/src/github.com/sunho/fws

ADD . .

RUN make

RUN mv out /fws

CMD ["/fws/server"]