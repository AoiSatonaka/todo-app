#!/bin/sh

proto_file_dir='.'
api_proto_files=$(find ${proto_file_dir} -type f -name '*.proto')
proto_out_dir='../docs'
proto_out_type='markdown'
proto_out_file='protobuf.md'

# change working directory
cd $(dirname $0)
cd ..

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
