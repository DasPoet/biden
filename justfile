set shell := ["powershell", "-ExecutionPolicy", "Bypass", "-c"]

fmt:
    go fmt ./...

test:
    go test ./...

bench-slices:
    cd benchmark; go test -bench="." -benchmem .

bench-slices-and-visualise:
    cd graphs; .venv\Scripts\activate; cd ../; just bench-slices | go run v1/internal/benchmark/cmd/main.go | python graphs/graphs.py
