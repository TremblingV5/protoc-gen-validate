module github.com/envoyproxy/protoc-gen-validate/tests

go 1.12

require (
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/golang/protobuf v1.4.2
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	google.golang.org/protobuf v1.25.0
)

replace github.com/envoyproxy/protoc-gen-validate => ../
