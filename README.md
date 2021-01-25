# go-skeleton

> A Golang package template repository. - 一个 Golang 软件包模板存储库。

[![tests](https://github.com/guanguans/go-skeleton/workflows/tests/badge.svg)](https://github.com/guanguans/go-skeleton/actions)
[![gocover.io](https://gocover.io/_badge/github.com/guanguans/go-skeleton)](https://gocover.io/github.com/guanguans/go-skeleton)
[![Go Report Card](https://goreportcard.com/badge/github.com/guanguans/go-skeleton)](https://goreportcard.com/report/github.com/guanguans/go-skeleton)
[![GoDoc](https://godoc.org/github.com/guanguans/go-skeleton?status.svg)](https://godoc.org/github.com/guanguans/go-skeleton)
[![GitHub release](https://img.shields.io/github/tag/guanguans/go-skeleton.svg)](https://github.com/guanguans/go-skeleton/releases)
[![GitHub license](https://img.shields.io/github/license/guanguans/go-skeleton.svg)](https://github.com/guanguans/go-skeleton/blob/master/LICENSE)

## Features

* package template

## Requirement

* Go >= 1.11

## Installation

``` shell script
$ go get -u github.com/guanguans/go-skeleton
```

## Usage

1. replace `guanguans/go-skeleton` -> `your github name/your package name`
2. replace `guanguans` -> `your github name`
2. replace `ityaozm@gmail.com` -> `your email`

This is just a quick introduction, view the [GoDoc](https://godoc.org/github.com/guanguans/go-skeleton) for details.

Let's start with a trivial example:

``` go
package main

import (
	"github.com/guanguans/skeleton"
	"gopkg.in/ffmt.v1"
)

func main() {
	ffmt.P(skeleton.ReturnTrue())
}
```

## Testing

``` bash
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
