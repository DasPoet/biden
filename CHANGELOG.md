# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- `biden` now supports UUIDs as defined in `github.com/google/uuid`
- `biden` now supports maps

## [0.2.0] - 2025-03-03

### Added

- benchmarks comparing `biden` to [benc](https://github.com/deneonet/benc) and [encoding/gob](https://pkg.go.dev/encoding/gob) have been added
- proper documentation has been added to the `README.md`

### Changed

- `biden` has been restructured to allow for easier consumption

## [0.1.0] - 2025-01-05

### Added

- `biden` now supports the following primitive types: `bool`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `int`, `int8`, `int16`, `int32`, `int64`, `float32`, `float64`, `string`
- `biden` now supports slices of the aforementioned primitive types
