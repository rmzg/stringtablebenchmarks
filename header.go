package main

import (
	"fmt"

	"github.com/alecthomas/mph"
)

func main() {
	//_ = Runes
	fmt.Printf("Test: %s\n", KV("UP ARROWHEAD BETWEEN TWO HORIZONTAL BARS"))
}

var chd *mph.CHD

func init() {
	var err error
	chd, err = mph.Mmap(data)
	if err != nil {
		panic(1)
	}
}

func KV(k string) []byte {
	return chd.Get([]byte(k))
}
