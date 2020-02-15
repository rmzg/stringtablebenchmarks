```
goos: linux
goarch: amd64
pkg: example.com/test/maptest
BenchmarkMap-48                      476           2281859 ns/op
BenchmarkMapNoInit-48            1007122              1189 ns/op
PASS
ok      example.com/test/maptest        3.471s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest
BenchmarkCHD-48              898           1397538 ns/op
PASS
ok      example.com/test/mphtest        1.395s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest2
BenchmarkMph2-48          986949              1219 ns/op
PASS
ok      example.com/test/mphtest2       1.946s
goos: linux
goarch: amd64
pkg: example.com/test/switchtest
BenchmarkSwitch-48        195211              5395 ns/op
PASS
ok      example.com/test/switchtest     1.121s
```
