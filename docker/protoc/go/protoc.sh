#!/bin/sh

proto_file_dir='./api'
proto_out_dir='./backend/internal/presentation/grpc'

# move to project root directory
cd $(dirname $0)
cd ../../../

# search protofiles
api_proto_files=$(find ${proto_file_dir} -type f -name '*.proto')

# out dir exists check
if [ ! -d $proto_out_dir ]; then
	mkdir -p $proto_out_dir
fi

# protobuf build
protoc \
	-I $proto_file_dir \
	--go_out=$proto_out_dir \
	--go_opt=paths=source_relative \
	--go-grpc_out=$proto_out_dir \
	--go-grpc_opt=paths=source_relative \
	$api_proto_files
