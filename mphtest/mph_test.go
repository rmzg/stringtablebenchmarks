package mphtest

import (
	"testing"

	"github.com/alecthomas/mph"
)

func BenchmarkCHD(b *testing.B) {
	chd = nil
	for i := 0; i < b.N; i++ {
		start()

		for j := 0; j < 20; j++ {
			x := KV("BLACK SMILING FACE")
			x = KV("SMILING FACE WITH OPEN MOUTH")
			x = KV("SMILING FACE WITH SUNGLASSES")
			x = KV("SMILING CAT FACE WITH OPEN MOUTH")
			x = KV("HAMMER AND SICKLE")
			x = KV("ROBOT FACE")
			_ = x

		}
	}

}

var chd *mph.CHD

func start() {
	var err error
	chd, err = mph.Mmap(data)
	if err != nil {
		panic(1)
	}
}

func KV(k string) []byte {
	return chd.Get([]byte(k))
}

