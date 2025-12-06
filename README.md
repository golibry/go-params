# go-params

A lightweight Go library for extracting and converting parameters from multiple sources with built‑in type conversion and sensible defaults.  
Migrated from https://github.com/rsgcata/go-params

## Installation

Add the dependency to your project:

```shell
go get github.com/golibry/go-params
```

## Features

- Environment variables parsing with type conversion
- URL query parameters parsing with type conversion
- Raw string conversion helpers
- Built‑in support for common types: `string`, `int`, `bool`, `float64`, `time.Duration`
- Consistent API across parameter sources
- Default values used automatically for missing or invalid inputs (no errors returned)

## Examples

See runnable examples in the `_examples` folder:

- `_examples/env_examples.go` — getting env vars as typed values and defaults
- `_examples/url_examples.go` — reading URL query parameters via `QueryParams`
- `_examples/strconv_examples.go` — converting raw strings to typed values

Clone the repo and run the files with `go run` to try them out.
