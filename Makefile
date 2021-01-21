build_path := build/bin
binary := image-transform

binfile := $(build_path)/$(binary)
repo := "github.com/image-transform"

.PHONY: build
build:
	@go build -race -gcflags='-N -l' -o "$(binfile)" "$(repo)/cmd/image-transform" 2>&1
run: build
	@$(binfile)

debug: build
	@echo "[!] Starting debug process"
	@echo -e "[*] Debug target: \e[33m${build_path}/${binary}\e[0m"
	@echo "**********************************************************"
	@dlv exec "${build_path}/${binary}"

debug-test:
	@go clean -testcache
	@go test -tags=unit `go list ./... | grep -i "${package}"` -gcflags='-l -N' -v -c -o "${build_path}/${binary}.test"
	dlv exec "${build_path}/${binary}.test" -- -test.run "${test_name}"

it:
	@go test -tags=integration -count=1 -v `go list ./...`

ut:
	@go test -tags=unit -count=1 -v `go list ./...`

ut-sub:
	@go test -run "${name}" -v -tags=unit -count=1 `go list ./...`

gen_mocks:
	mockery -r --dir ./pkg --name FS --output internal/mocks --outpkg mocks --case underscore
