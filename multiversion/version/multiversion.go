package main

import (
	"github.com/hase-yuu/module-sample/multiversion/a"
	"github.com/hase-yuu/module-sample/multiversion/b"
)

func main() {
	a.Call()
	b.Call()
}