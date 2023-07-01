
buf-go-clean:
	rm -rf base/generate
buf-base:
	cd base && go mod tidy
	cd base && go get github.com/grpc-ecosystem/grpc-gateway/v2
	cd base && go get google.golang.org/genproto/googleapis/api/annotations
buf-go:
	cd base/proto && buf generate 
buf-update:
	cd base/proto && buf mod update
buf-format:
	cd base/proto && buf format -w -d
buf-lint:
	cd base/proto && buf lint
buf-doc:
	protoc --doc_out=./base/docs --doc_opt=html,index.html ./base/proto/*/*/*.proto 

buf-gen: buf-base buf-update buf-format buf-lint buf-go-clean buf-go buf-doc