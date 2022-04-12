# DateTime - Simple DateTime Engine for Go

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/datetime)](https://pkg.go.dev/github.com/go-zoox/datetime)
[![Build Status](https://github.com/go-zoox/datetime/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/datetime/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/datetime)](https://goreportcard.com/report/github.com/go-zoox/datetime)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/datetime/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/datetime?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/datetime.svg)](https://github.com/go-zoox/datetime/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/datetime.svg?label=Release)](https://github.com/go-zoox/datetime/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/datetime
```

## Getting Started

```go
fmt.Println(datetime.Now().Format("YYYY-MM-DD HH:mm:ss"))
// 2022-04-12 13:05:50
```

## Inspired by
* [zcorky/moment](https://github.com/zcorky/moment) - A minimalist JavaScript library that parses, validates, manipulates, and displays dates and times.
* dayjs
* moment.js

## License
GoZoox is released under the [MIT License](./LICENSE).