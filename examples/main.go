package main

import (
	"github.com/guanguans/skeleton"
	"github.com/spf13/cast"
	"gopkg.in/ffmt.v1"
)

func main() {
	// ffmt
	ffmt.P(skeleton.ReturnSelf("go-skeleton"))

	// spf13/cast
	ffmt.P(cast.ToInt("080"))
}
