# benchtest-arr-go
This code is a example to validate performance using append or not in golang

# result benchtests
```
go test -benchmem -bench .  
goos: darwin
goarch: arm64
pkg: poc
BenchmarkTestWithLength-8      	   30582	     38545 ns/op	  386297 B/op	      20 allocs/op
BenchmarkTestWithoutLength-8   	  129008	      9282 ns/op	   81920 B/op	       1 allocs/op
BenchmarkTestNotUseAppend-8    	  292729	      3985 ns/op	   81920 B/op	       1 allocs/op
PASS
ok  	poc	4.578s
```
