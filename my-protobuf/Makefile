# Change this to your own Go module
GO_MODULE := my-protobuf


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	del /s /q protogen 2>nul || exit 0
else
	rm -rf ./protogen
endif

.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \
	./proto/dummy/*.proto \
	./proto/jobsearch/*.proto \

.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run


.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
