#!/bin/sh

proto_file_dir='./api'
proto_out_dir='./docs'
proto_out_type='markdown'
proto_out_file='protobuf.md'

# move to project root directory
cd $(dirname $0)
cd ../../../

# search protofiles
api_proto_files=$(find ${proto_file_dir} -type f -name '*.proto')
pwd

# out dir exists check
if [ ! -d $proto_out_dir ]; then
	mkdir $proto_out_dir
fi

# protobuf build
protoc \
	-I $proto_file_dir \
	--doc_out=$proto_out_dir \
	--doc_opt="${proto_out_type},${proto_out_file}" \
	$api_proto_files
