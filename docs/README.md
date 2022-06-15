# go-skeleton

[ENGLISH](README.md) | [简体中文](README-zh_CN.md)

> A Go package template repository. - 一个 Go 软件包模板存储库。

[![tests](https://github.com/guanguans/go-skeleton/workflows/tests/badge.svg)](https://github.com/guanguans/go-skeleton/actions)
[![gocover.io](https://gocover.io/_badge/github.com/guanguans/go-skeleton)](https://gocover.io/github.com/guanguans/go-skeleton)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/go-skeleton)](https://goreportcard.com/report/github.com/guanguans/go-skeleton)
[![GoDoc](https://pkg.go.dev/github.com/guanguans/go-skeleton?status.svg)](https://pkg.go.dev/github.com/guanguans/go-skeleton)
[![GitHub license](https://img.shields.io/github/license/guanguans/go-skeleton.svg)](https://github.com/guanguans/go-skeleton/blob/master/LICENSE)
[![Open in Visual Studio Code](https://open.vscode.dev/badges/open-in-vscode.svg)](https://open.vscode.dev/guanguans/go-skeleton)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/guanguans/go-skeleton)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/guanguans/go-skeleton)
![GitHub repo size](https://img.shields.io/github/repo-size/guanguans/go-skeleton)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/guanguans/go-skeleton)

## Features

* package template

## Requirement

* Go >= 1.11

## Installation

```shell script
$ go get -u github.com/guanguans/go-skeleton
```

## Usage

1. replace `guanguans/go-skeleton` -> `your github name/your repository name`
2. replace `go-skeleton` -> `your repository name`
3. replace `skeleton` -> `your package name`
4. replace `guanguans` -> `your github name`
5. replace `ityaozm@gmail.com` -> `your email`

This is just a quick introduction, view the [GoDoc](https://pkg.go.dev/github.com/guanguans/go-skeleton) for details.

Let's start with a trivial example:

```go
package main

import (
	"github.com/guanguans/go-skeleton"
	"gopkg.in/ffmt.v1"
)

func main() {
	ffmt.P(skeleton.ReturnSelf("go-skeleton"))
}
```

## Testing

```shell script
$ make test
$ make bench
```

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

* [guanguans](https://github.com/guanguans)
* [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
