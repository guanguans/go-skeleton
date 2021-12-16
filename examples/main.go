package main

import (
	"github.com/awesee/php2go/php"
	"github.com/guanguans/skeleton"
	"github.com/spf13/cast"
	"github.com/syyongx/php2go"
	"gopkg.in/ffmt.v1"
)

func main() {
	// ffmt
	ffmt.P(skeleton.ReturnSelf("go-skeleton"))

	// spf13/cast
	ffmt.P(cast.ToInt("80"))

	// syyongx/php2go
	ffmt.P(php.Md5("golang"))

	// awesee/php2go
	ffmt.P(php2go.Md5("golang"))
}
