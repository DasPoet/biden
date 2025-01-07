fmt:
    go fmt ./...

test:
    go test ./...

bench-slices:
    go test -bench="." -benchmem slices_benchmark_test.go helpers_test.go primitives.go slices.go biden.go
