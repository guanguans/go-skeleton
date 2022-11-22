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
![GitHub all releases](https://img.shields.io/github/downloads/guanguans/go-skeleton/total)

## 参考项目

* [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
* [golang-repo-template](https://github.com/moul/golang-repo-template)

## 功能

* package template

## 环境要求

* Go >= 1.11

## 安装

```shell script
$ go get -u github.com/guanguans/go-skeleton
```

## 使用

1. replace `guanguans/go-skeleton` -> `your github name/your repository name`
2. replace `go-skeleton` -> `your repository name`
3. replace `skeleton` -> `your package name`
4. replace `guanguans` -> `your github name`
5. replace `ityaozm@gmail.com` -> `your email`

这只是一个快速介绍, 请查看 [GoDoc](https://pkg.go.dev/github.com/guanguans/go-skeleton) 获得详细信息。

让我们从一个简单的例子开始：

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

## 测试

```shell script
$ make test
$ make bench
```

## 变更日志

请参阅 [CHANGELOG](CHANGELOG.md) 获取最近有关更改的更多信息。

## 贡献指南

请参阅 [CONTRIBUTING](.github/CONTRIBUTING.md) 有关详细信息。

## 安全漏洞

请查看[我们的安全政策](../../security/policy)了解如何报告安全漏洞。

## 贡献者

* [guanguans](https://github.com/guanguans)
* [所有贡献者](../../contributors)

## 协议

MIT 许可证（MIT）。有关更多信息，请参见[协议文件](LICENSE)。
