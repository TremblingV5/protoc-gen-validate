empty :=
space := $(empty) $(empty)
PACKAGE := github.com/envoyproxy/protoc-gen-validate

# protoc-gen-go parameters for properly generating the import path for PGV
VALIDATE_IMPORT := Mvalidate/validate.proto=${PACKAGE}/validate
GO_IMPORT_SPACES := ${VALIDATE_IMPORT},\
	Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,\
	Mgoogle/protobuf/duration.proto=github.com/golang/protobuf/ptypes/duration,\
	Mgoogle/protobuf/struct.proto=github.com/golang/protobuf/ptypes/struct,\
	Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
	Mgoogle/protobuf/wrappers.proto=github.com/golang/protobuf/ptypes/wrappers,\
	Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor
GO_IMPORT:=$(subst $(space),,$(GO_IMPORT_SPACES))

PROTOC_GEN_VALIDATE_LOCALES_DIR := /go/src/github.com/envoyproxy/protoc-gen-validate/tests/harness/locales
export PROTOC_GEN_VALIDATE_LOCALES_DIR

.PHONY: build
build: validate/validate.pb.go
	# generates the PGV binary and installs it into $$GOPATH/bin
	go install .

.PHONY: bazel
bazel:
	# generate the PGV plugin with Bazel
	bazel build //tests/...

.PHONY: build_generation_tests
build_generation_tests:
	bazel build //tests/generation/...

.PHONY: gazelle
gazelle:
	# runs gazelle against the codebase to generate Bazel BUILD files
	bazel run //:gazelle -- update-repos -from_file=go.mod -prune -to_macro=dependencies.bzl%go_third_party
	bazel run //:gazelle
	buildozer 'replace deps @com_github_golang_protobuf//ptypes:go_default_library_gen @com_github_golang_protobuf//ptypes:go_default_library' '//...:%go_library'
	buildozer 'replace deps @io_bazel_rules_go//proto/wkt:duration_go_proto @com_github_golang_protobuf//ptypes/duration:go_default_library' '//...:%go_library'
	buildozer 'replace deps @io_bazel_rules_go//proto/wkt:timestamp_go_proto @com_github_golang_protobuf//ptypes/timestamp:go_default_library' '//...:%go_library'

.PHONY: lint
lint: bin/golint bin/shadow
	# lints the package for common code smells
	test -z "$(gofmt -d -s ./*.go)" || (gofmt -d -s ./*.go && exit 1)
	# golint -set_exit_status
	# check for variable shadowing
	go vet -vettool=$(shell pwd)/bin/shadow ./...

bin/shadow:
	GOBIN=$(shell pwd)/bin go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

bin/golint:
	GOBIN=$(shell pwd)/bin go install golang.org/x/lint/golint

bin/protoc-gen-go:
	GOBIN=$(shell pwd)/bin go install github.com/golang/protobuf/protoc-gen-go

bin/harness:
	cd tests && go build -o ../bin/harness ./harness/executor

.PHONY: harness
harness: testcases tests/harness/go/harness.pb.go tests/harness/go/main/go-harness tests/harness/cc/cc-harness bin/harness
 	# runs the test harness, validating a series of test cases in all supported languages
	./bin/harness -go -cc

.PHONY: bazel-harness
bazel-harness:
	# runs the test harness via bazel
	bazel run //tests/harness/executor:executor --incompatible_new_actions_api=false -- -go -cc -java -python

.PHONY: testcases
testcases: bin/protoc-gen-go
	# generate the test harness case protos
	rm -r tests/harness/cases/go || true
	mkdir tests/harness/cases/go
	rm -r tests/harness/cases/other_package/go || true
	mkdir tests/harness/cases/other_package/go
	# copy test cases for custom error messages
	rm -r tests/harness/cases_custom || true
	cp -r tests/harness/cases tests/harness/cases_custom && \
		cd tests/harness/cases_custom && \
		sed -i'' -e 's/package tests.harness.cases;/package tests.harness.cases_custom;/' *.proto && \
		sed -i'' -e 's/option go_package = "cases";/option go_package = "cases_custom";/' *.proto && \
		sed -i'' -e 's/import "tests\/harness\/cases\/other_package\/embed.proto";/import "tests\/harness\/cases_custom\/other_package_custom\/other_embed.proto";/' *.proto && \
		sed -i'' -e 's/tests.harness.cases.other_package.Embed/tests.harness.cases_custom.other_package_custom.Embed/' *.proto && \
		sed -i'' -e 's/other_package.Embed/other_package_custom.Embed/' *.proto && \
		for filename in *.proto; do mv "$$filename" "other_$$filename"; done
	mv tests/harness/cases_custom/other_package tests/harness/cases_custom/other_package_custom && \
		cd tests/harness/cases_custom/other_package_custom && \
		sed -i'' -e 's/option go_package = "other_package";/option go_package = "other_package_custom";/' *.proto && \
		sed -i'' -e 's/package tests.harness.cases.other_package;/package tests.harness.cases_custom.other_package_custom;/' *.proto && \
		for filename in *.proto; do mv "$$filename" "other_$$filename"; done
	# protoc-gen-go makes us go a package at a time
	cd tests/harness/cases/other_package && \
	protoc \
		-I . \
		-I ../../../.. \
		--go_out="${GO_IMPORT}:./go" \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--validate_out="lang=go:./go" \
		./*.proto
	cd tests/harness/cases && \
	protoc \
		-I . \
		-I ../../.. \
		--go_out="Mtests/harness/cases/other_package/embed.proto=${PACKAGE}/tests/harness/cases/other_package/go,${GO_IMPORT}:./go" \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--validate_out="lang=go,Mtests/harness/cases/other_package/embed.proto=${PACKAGE}/tests/harness/cases/other_package/go:./go" \
		./*.proto
	# generate for custom error messages
	cd tests/harness/cases_custom/other_package_custom && \
	protoc \
		-I . \
		-I ../../../.. \
		--go_out="${GO_IMPORT}:./go" \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--validate_out="lang=go:./go" \
		./*.proto && \
	cd go && sed -i'' -e 's/	\(.*\) "tests\/harness\/\(.*\)"/\1 "github.com\/envoyproxy\/protoc-gen-validate\/tests\/harness\/\2\/go"/' *.go
	cd tests/harness/cases_custom && \
	protoc \
		-I . \
		-I ../../.. \
		--go_out="Mtests/harness/cases_custom/other_package_custom/embed.proto=${PACKAGE}/tests/harness/cases_custom/other_package_custom/go,${GO_IMPORT}:./go" \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--validate_out="lang=go,Mtests/harness/cases_custom/other_package_custom/embed.proto=${PACKAGE}/tests/harness/cases_custom/other_package_custom/go:./go" \
		./*.proto && \
	cd go && sed -i'' -e 's/	\(.*\) "tests\/harness\/\(.*\)"/\1 "github.com\/envoyproxy\/protoc-gen-validate\/tests\/harness\/\2\/go"/' *.go

validate/validate.pb.go: bin/protoc-gen-go validate/validate.proto
	cd validate && protoc -I . \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--go_opt=paths=source_relative \
		--go_out="${GO_IMPORT}:." validate.proto

tests/harness/go/harness.pb.go: bin/protoc-gen-go tests/harness/harness.proto
	# generates the test harness protos
	cd tests/harness && protoc -I . \
		--plugin=protoc-gen-go=$(shell pwd)/bin/protoc-gen-go \
		--go_out="${GO_IMPORT}:./go" harness.proto

tests/harness/go/main/go-harness:
	# generates the go-specific test harness
	cd tests && go build -o ./harness/go/main/go-harness ./harness/go/main

tests/harness/cc/cc-harness: tests/harness/cc/harness.cc
	# generates the C++-specific test harness
	# use bazel which knows how to pull in the C++ common proto libraries
	bazel build //tests/harness/cc:cc-harness
	cp bazel-bin/tests/harness/cc/cc-harness $@
	chmod 0755 $@

tests/harness/java/java-harness:
	# generates the Java-specific test harness
	mvn -q -f java/pom.xml clean package -DskipTests

.PHONY: ci
ci: lint bazel testcases bazel-harness build_generation_tests

.PHONY: clean
clean:
	(which bazel && bazel clean) || true
	rm -f \
		bin/protoc-gen-go \
		bin/harness \
		tests/harness/cc/cc-harness \
		tests/harness/go/main/go-harness \
		tests/harness/go/harness.pb.go
	rm -rf \
		tests/harness/cases/go \
		tests/harness/cases/other_package/go


