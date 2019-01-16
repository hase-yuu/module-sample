package main // import "github.com/hase-yuu/module-sample/directive"

import (
	"fmt"
	v1 "rsc.io/quote"
	v2 "rsc.io/quote/v2"
	v3 "rsc.io/quote/v3"
)


func main() {
	fmt.Println(v1.Hello())
	fmt.Println(v2.HelloV2())
	fmt.Println(v3.HelloV3())
}
