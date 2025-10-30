.PHONY: build
build:
	go build -gcflags "all=-N -l"

.PHONY: test
test:
	go test -timeout 3m ./...

.PHONY: lint
lint:
	@golangci-lint run --timeout 10m

.PHONY: gen-docs
gen-docs: build
	rm -rf ./docs/tables/*
	mkdir -p ./docs/tables
	# Use cloudquery command from PATH to generate docs
	cloudquery tables docs/spec.yml --output-dir . --format markdown

# All gen targets
.PHONY: gen
gen: gen-docs

.PHONY: serve
serve:
	AWS_PROFILE=deployTools go run main.go serve
