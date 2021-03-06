FROM golang:1.11.4-stretch

RUN apt-get update && apt-get install -y --no-install-recommends \
    autoconf \
    automake \
    libtool

ENV PROTOBUF_VERSION 3.6.1
RUN wget -O protobuf.tar.gz "https://github.com/google/protobuf/releases/download/v${PROTOBUF_VERSION}/protobuf-cpp-${PROTOBUF_VERSION}.tar.gz"
RUN tar xzf protobuf.tar.gz && cd "protobuf-${PROTOBUF_VERSION}" && ./autogen.sh && ./configure && make && make install && ldconfig

ENV GO_PROTOBUF_VERSION 1.2.0
RUN git clone https://github.com/golang/protobuf /go/src/github.com/golang/protobuf && cd /go/src/github.com/golang/protobuf && git checkout "v${GO_PROTOBUF_VERSION}"
RUN cd /go/src/github.com/golang/protobuf/protoc-gen-go && go get ./... && go install github.com/golang/protobuf/protoc-gen-go

ENV GRPC_GATEWAY_VERSION 1.6.3
RUN git clone https://github.com/grpc-ecosystem/grpc-gateway /go/src/github.com/grpc-ecosystem/grpc-gateway && cd /go/src/github.com/grpc-ecosystem/grpc-gateway && git checkout "v${GRPC_GATEWAY_VERSION}"
RUN cd /go/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && go get ./... && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN cd /go/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && go get ./... && go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

RUN rm -rf /var/lib/apt/lists/* protobuf.tar.gz
