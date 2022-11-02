// This file is part of the guanguans/go-skeleton.
// (c) 2021. guanguans <ityaozm@gmail.com>
// This source file is subject to the MIT license that is bundled.

package main

import (
	"github.com/awesee/php2go/php"
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/go-skeleton"
	"github.com/kr/pretty"
	"github.com/spf13/cast"
	"github.com/syyongx/php2go"
	"github.com/zakaria-chahboun/cute"
	"gopkg.in/ffmt.v1"
)

func main() {
	// ffmt
	ffmt.P(skeleton.ReturnSelf("go-skeleton"))

	// davecgh/go-spew
	spew.Dump("go-skeleton")

	// spf13/cast
	ffmt.P(cast.ToInt("go-skeleton"))

	// syyongx/php2go
	ffmt.P(php.Md5("go-skeleton"))

	// awesee/php2go
	ffmt.P(php2go.Md5("go-skeleton"))

	// zakaria-chahboun/cute
	cute.Println("go-skeleton")

	// kr/pretty
	pretty.Println("go-skeleton")
}
