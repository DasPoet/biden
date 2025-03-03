$$\LARGE {\color{white}\textrm{bi}}{\color{gray}\textrm{nary~}}{\color{white}\textrm{d}}{\color{gray}\textrm{ata~}}{\color{white}\textrm{en}}{\color{gray}\textrm{coder}} $$

<div align="center">
    <img src="https://img.shields.io/badge/Written_In-Go-00acd7?style=for-the-badge&logo=go" alt="Go" style="height: 28px;" />
    &nbsp;&nbsp;
    <img src="https://github.com/DasPoet/biden/actions/workflows/go.yml/badge.svg?branch=master" alt="Pipeline status" style="height: 28px;" />
</div>

<div align="center">
    <img width="300" src="/assets/logo.png" alt="logo" />
</div>

<div align="center">
    Biden *(baɪ̯dən) * provides low-level facilities for converting between Go's primitive types and binary data streams.
</div

<hr />

## 📦 Getting started

### Installation

```bash
go get -u github.com/DasPoet/biden
```

### Usage

```go
package main

import "github.com/DasPoet/biden/pkg/biden"

func main() {
    // TODO
}
```

## 🚀 Benchmarks

There are several libraries that cover the same usecase as `biden`. `biden` was written to offer the best possible performance. To test whether it is any faster than competing alternatives, this repository contains a number of benchmarks. Their results are displayed below. If you are choosing this library for reasons related to performance, be sure to verify the results of these benchmarks in your own execution environment. Don't just believe things people are saying on the internet. If you find any mistakes in any of the code in this repository (including the benchmarks), please open an issue.

The benchmarks compare `biden` to:

* [benc](https://github.com/deneonet/benc)
* [encoding/gob](https://pkg.go.dev/encoding/gob)

If you feel like we should include any other repository in this list, please suggest an addition by opening an issue.

### Marshaling

![marshal slices](assets/benchmark_marshal.svg)

### Unmarshaling

![unmarshal slices](assets/benchmark_unmarshal.svg)

