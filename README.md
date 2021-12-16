# benchtest-arr-go
This code is a example to validate performance using append or not in golang

# result benchtests
```
go test -benchmem -bench .
goos: darwin
goarch: arm64
pkg: poc
BenchmarkTestWithLength-8         115262              9422 ns/op           81920 B/op          1 allocs/op
BenchmarkTestWithoutLength-8       30625             39379 ns/op          386298 B/op         20 allocs/op
BenchmarkTestNotUseAppend-8       285175              4103 ns/op           81920 B/op          1 allocs/op
PASS
ok      poc     4.218s
```
