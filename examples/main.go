package main

import (
	"github.com/awesee/php2go/php"
	"github.com/davecgh/go-spew/spew"
	"github.com/guanguans/skeleton"
	"github.com/spf13/cast"
	"github.com/syyongx/php2go"
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
}
