FROM golang:1.21.2
ARG PROTOBUF_VERSION=24.4
ARG PROTOC_GEN_GO_VERSION=1.31
ARG PROTOC_GEN_GO_GRPC_VERSION=1.3.0

# unzip install
RUN apt update && apt install unzip

# protoc install
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip &&\
  unzip protoc.zip && \
  mv bin/* /usr/local/bin/ && \
  mv include/* /usr/local/include/

# protoc plugin install
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}
