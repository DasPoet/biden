fmt:
    go fmt ./...

test:
    go test ./...

bench-slices:
    cd pkg/benchmark && go test -bench="." -benchmem .

bench-slices-and-visualise:
    cd graphs && source .venv/bin/activate && cd .. && just bench-slices | go run internal/benchmark/cmd/main.go | python graphs/graphs.py
