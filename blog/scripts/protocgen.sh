#!/usr/bin/env bash

set -e

echo "Generating gogo proto code:"
cd proto
proto_dirs=$(find . -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
    echo "proto_dirs = $proto_dirs"
for dir in $proto_dirs; do
  for file in $(find "${dir}" -maxdepth 1 -name '*.proto'); do
    # this regex checks if a proto file has its go_package set to blog/cosmossdk/api/api/...
    # gogo proto files SHOULD ONLY be generated if this is false
    # we don't want gogo proto to run for proto files which are natively built for google.golang.org/protobuf
    echo "Generating gogo proto code for $file"
    if grep -q "option go_package" "$file" && grep -H -o -c 'option go_package.*blog/cosmossdk/api' "$file" | grep -q ':0$'; then
      buf generate --template buf.gen.gogo.yaml $file
    fi
  done
done

echo "Generating pulsar proto code"
buf generate --template buf.gen.pulsar.yaml

cd ..


# Move generated files to the correct directories
rm -rf api && mkdir api
mv proto/api/* ./api
mv proto/blog/x/blog/types/* ./x/blog/types

# Remove proto directories
rm -rf proto/api proto/blog/x