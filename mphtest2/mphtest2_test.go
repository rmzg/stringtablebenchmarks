package mphtest2

import "testing"

func BenchmarkMph2(b *testing.B) {
	for i := 0; i < b.N; i++ {

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
