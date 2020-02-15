```
goos: linux
goarch: amd64
pkg: example.com/test/maptest
BenchmarkMap-48                      460           2368314 ns/op
BenchmarkMapNoInit-48             992005              1213 ns/op
PASS
ok      example.com/test/maptest        3.474s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest
BenchmarkCHD-48              890           1351237 ns/op
PASS
ok      example.com/test/mphtest        1.345s
goos: linux
goarch: amd64
pkg: example.com/test/mphtest2
BenchmarkMph2-48                    3242            360799 ns/op
BenchmarkMph2NoIinit-48          1003911              1185 ns/op
PASS
ok      example.com/test/mphtest2       3.196s
goos: linux
goarch: amd64
pkg: example.com/test/switchtest
BenchmarkSwitch-48        186889              5463 ns/op
PASS
ok      example.com/test/switchtest     1.091s
```
