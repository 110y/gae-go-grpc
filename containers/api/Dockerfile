FROM golang:1.11.4-stretch

ENV GO_REALIZE_VERSION=2.0.2

RUN git clone https://github.com/oxequa/realize.git /go/src/github.com/oxequa/realize && cd /go/src/github.com/oxequa/realize && git checkout "v${GO_REALIZE_VERSION}"
RUN cd /go/src/github.com/oxequa/realize && go get ./... && go install github.com/oxequa/realize
