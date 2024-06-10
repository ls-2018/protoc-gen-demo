gen_example:
	go install
	protoc -I ./example/api \
 	--go_out ./example/api --go_opt=paths=source_relative \
	--go-gin_out ./example/api --go-gin_opt=paths=source_relative \
	example/api/product/app/v1/v1.proto


install:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


debug:
	mkdir debug
	go mod vendor
	#gsed -i 's@import (@import ( \n "time" @g' vendor/google.golang.org/protobuf/compiler/protogen/protogen.go
	gsed -i 's@pluginpb.CodeGeneratorRequest{}@pluginpb.CodeGeneratorRequest{} \n ioutil.WriteFile("debug/tmp.txt", in, os.ModePerm) @g' vendor/google.golang.org/protobuf/compiler/protogen/protogen.go
	go fmt vendor/google.golang.org/protobuf/compiler/protogen/protogen.go

