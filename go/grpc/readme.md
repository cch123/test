brew install protoc-gen-go

protoc --proto_path=./ --go_out=plugins=grpc:./ ./hello.proto
