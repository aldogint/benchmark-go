```zsh
$ go test -bench=. -benchtime=12604x                                                                                                                                                                         
goos: linux                                                                                                                                                                                                         
goarch: amd64                                                                                                                                                                                                       
pkg: github.com/aldogint/benchmark                                                                                                                                                                                  
cpu: Intel(R) Core(TM) i7-10710U CPU @ 1.10GHz                                                                                                                                                                      
BenchmarkRangeNode-12              12604            138960 ns/op                                                                                                                                                    
BenchmarkRangeBlock-12             12604               496.2 ns/op 
```
