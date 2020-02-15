goos: linux
goarch: amd64
pkg: example.com/test/maptest
BenchmarkMap-48              471           2361129 ns/op
PASS
ok      example.com/test/maptest        1.376s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest
BenchmarkCHD-48              840           1388397 ns/op
PASS
ok      example.com/test/mphtest        1.316s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest2
BenchmarkMph2-48          992220              1208 ns/op
PASS
ok      example.com/test/mphtest2       1.986s
goos: linux
goarch: amd64
pkg: example.com/test/switchtest
BenchmarkSwitch-48        200536              5368 ns/op
PASS
ok      example.com/test/switchtest     1.143s
