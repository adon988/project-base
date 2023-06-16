
buf-go-clean:
	rm -rf base/generate/proro
buf-go:
	cd base && buf generate && go mod tidy
	cd base && go get github.com/grpc-ecosystem/grpc-gateway/v2
buf-format:
	cd base && buf format -w -d
buf-lint:
	cd base && buf lint
buf-doc:
	protoc --doc_out=./base/docs --doc_opt=html,index.html ./base/proto/*/*/*.proto 

buf-gen: buf-format buf-lint buf-go-clean buf-go buf-doc