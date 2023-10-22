FROM node:20-slim
ARG PROTOBUF_VERSION=24.4
ARG PROTOC_GEN_ELM_VERSION=3.2.0

# unzip install
RUN apt update && apt install -y curl unzip

# protoc install
WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip &&\
  unzip protoc.zip && \
  mv bin/* /usr/local/bin/ && \
  mv include/* /usr/local/include/

# protoc plugin install
RUN npm i -g protoc-gen-elm@${PROTOC_GEN_ELM_VERSION}
