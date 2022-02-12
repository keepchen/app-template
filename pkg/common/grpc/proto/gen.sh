#!/bin/sh
PROTO_FILES=""

WORKDIR="pkg/common/grpc/proto/"

# shellcheck disable=SC2164
cd ${WORKDIR}

echo "|--proto file(s) list:--|"
for file in *.proto
do
    echo "$file"
    if test -f "$file"
    then
      # shellcheck disable=SC2039
      if [ -z "$PROTO_FILES" ];then
        PROTO_FILES="$file"
      else
        PROTO_FILES="$PROTO_FILES $file" # 追加拼凑
      fi
    fi
done

protoc -I . $PROTO_FILES --go_out=plugins=grpc:../pb/