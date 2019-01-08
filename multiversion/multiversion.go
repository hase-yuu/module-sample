package main // import "github.com/hase-yuu/vgo-sample/multiversion

import (
	"github.com/hase-yuu/vgo-sample/multiversion/v1"
	"github.com/hase-yuu/vgo-sample/multiversion/v2"
)

func main() {
	v1.Call()
	v2.Call()
}