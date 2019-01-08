package main

import (
	"github.com/hase-yuu/module-sample/multiversion/a"
	"github.com/hase-yuu/module-sample/multiversion/b"
	"github.com/hase-yuu/module-sample/multiversion/c"
)

func main() {
	a.Call()
	b.Call()
	c.Call()
}