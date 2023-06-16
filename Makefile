
buf-go-clean:
	rm -rf base/generate/proro
buf-go:
	cd base && buf generate && go mod tidy
buf-format:
	cd base && buf format -w -d
buf-lint:
	cd base && buf lint

buf-gen: buf-format buf-lint buf-go-clean buf-go 