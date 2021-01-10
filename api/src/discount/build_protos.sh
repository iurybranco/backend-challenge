PROTO_DIR_NAME=./src/discount

# Generate JavaScript code
yarn run grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:${PROTO_DIR_NAME} \
    --grpc_out=grpc_js:${PROTO_DIR_NAME} \
    --plugin=protoc-gen-grpc=./node_modules/.bin/grpc_tools_node_protoc_plugin \
    -I src/discount \
    proto/*.proto

# Generate TypeScript code
yarn run grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
    --ts_out=grpc_js:${PROTO_DIR_NAME} \
    -I src/discount \
    proto/*.proto
