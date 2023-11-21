FROM golang:1.21.2
ARG PROTOBUF_VERSION=24.4
ARG PROTOC_GEN_DOC_VERSION=1.5.1

# unzip install
RUN apt update && apt install unzip

# protoc install
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip &&\
  unzip protoc.zip && \
  mv bin/* /usr/local/bin/ && \
  mv include/* /usr/local/include/

# protoc plugin install
RUN go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v${PROTOC_GEN_DOC_VERSION}
