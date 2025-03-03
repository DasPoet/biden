package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/tools/benchmark/parse"
)

type BenchmarkType string

const (
	MarshalBenchmark   BenchmarkType = "marshal"
	UnmarshalBenchmark BenchmarkType = "unmarshal"
)

type BenchmarkResult struct {
	Library    string        `json:"benchmarkLibrary"`
	Type       BenchmarkType `json:"benchmarkType"`
	DataType   string        `json:"dataType"`
	NanosPerOp int           `json:"nanosPerOp"`
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var (
		scanner    = bufio.NewScanner(os.Stdin)
		benchmarks []BenchmarkResult
	)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}

		bench, err := parse.ParseLine(line)
		if err != nil {
			return err
		}

		parsed, err := parseBenchmark(*bench)
		if err != nil {
			return err
		}
		benchmarks = append(benchmarks, *parsed)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	encoded, err := json.MarshalIndent(benchmarks, "", " ")
	if err != nil {
		return err
	}

	fmt.Println(string(encoded))
	return nil
}

func parseBenchmark(raw parse.Benchmark) (*BenchmarkResult, error) {
	name, err := parseBenchmarkName(raw.Name)
	if err != nil {
		return nil, err
	}

	parsed := BenchmarkResult{
		Library:    name.Library,
		Type:       name.Type,
		DataType:   name.DataType,
		NanosPerOp: int(raw.NsPerOp),
	}
	return &parsed, nil
}

type ParsedBenchmarkName struct {
	Library  string
	Type     BenchmarkType
	DataType string
}

func parseBenchmarkName(name string) (*ParsedBenchmarkName, error) {
	var (
		pattern = regexp.MustCompile(`Benchmark_(?<lib>.*)_(?<type>Marshal|Unmarshal)(?<dataType>.*)-\d+`)
		groups  = extractNamedCapturingGroups(pattern, name)
	)

	var typ BenchmarkType
	switch groups["type"] {
	case "Marshal":
		typ = MarshalBenchmark
	case "Unmarshal":
		typ = UnmarshalBenchmark
	default:
		return nil, errors.New("invalid benchmark name: " + name)
	}

	parsed := &ParsedBenchmarkName{
		Library:  groups["lib"],
		Type:     typ,
		DataType: groups["dataType"],
	}
	return parsed, nil
}

func extractNamedCapturingGroups(pattern *regexp.Regexp, target string) map[string]string {
	var (
		matches = pattern.FindStringSubmatch(target)
		groups  = make(map[string]string)
	)

	for i, name := range pattern.SubexpNames() {
		if i > 0 && name != "" {
			groups[name] = matches[i]
		}
	}
	return groups
}
