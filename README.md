dla
===

Go implementation of DLA

    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 1 1 1 . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . 1 1 1 . . . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 1 1 1 1 1 . . . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . 1 . 1 1 . . . . . . . . . . 1 . . . . 1 1 1 1 1 . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . 1 1 1 1 1 . 1 1 . . 1 1 2 1 1 1 . . . 1 1 . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . 1 1 1 1 . . . . . . . 2 . . . . 1 . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 1 . . 1 1 2 2 . . . . . . . . . . 1 . . . . 1 1 1 . . . . . . . . . . . . . . . . 
    . . . . . . . . . 1 . . . 1 1 . . . . 1 . . . . . . . . . . 1 2 . 1 1 2 1 1 . . . . . . . . . . . . . . . . 
    . . . . . . . . 1 1 1 1 1 1 1 1 1 1 1 2 . . . . . . . 1 . . 1 1 1 2 . . 1 1 . . . . . . . . . . . . . . . . 
    . . . . . . . . . . 1 1 . 1 . 1 . . . 1 . 1 . . . . . 1 2 1 1 . . 1 1 . . 1 . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . 1 . 2 1 . . . . . . 1 . . 1 . 1 . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . 1 . 2 . . . . 1 1 . 1 . 1 1 1 1 . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . 1 . 1 . 1 2 . 1 . . . . . 1 1 1 1 1 1 . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . 1 1 1 1 1 1 1 1 1 1 1 1 1 1 . . . . 2 . 1 . 1 . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . 1 1 . 1 1 . . . . . . 1 1 . . . . . 2 1 3 . . 1 1 . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 1 1 . . . 1 1 1 1 . . . . 2 1 . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 1 . . . . 2 . . 1 1 . 1 2 1 1 . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . 1 1 1 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . 1 1 . 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . 1 1 2 . . . . . . 1 1 1 1 1 . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . 1 1 2 1 1 2 1 . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . 1 . . 1 1 . 1 . . 1 2 1 1 1 . . . . . . 1 . . . 1 . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . 1 1 1 . . 1 2 1 1 1 1 1 . 1 1 . . 1 . . 1 . 1 . 1 . 1 . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . 2 1 2 1 1 . . . 1 . 1 1 1 . 1 1 . 1 . 1 . 1 . 1 . . . . . . . . 
    . . . . . . . . . . . . . . . . . . 1 . . . . 1 1 . . . . . 1 . . 1 2 1 1 1 1 1 1 1 1 1 2 1 . 1 1 1 1 . . . 
    . . . . . . . . . . . . . . . . . . 1 . 1 . . . 1 2 1 1 . . 2 2 1 . 1 . . . 1 . . 1 . 1 . 1 1 1 1 . . . . . 
    . . . . . . . . . . . . . . 1 1 1 1 1 1 1 1 1 1 1 1 . . . . . . . . 1 . . . . . . 1 1 . . 1 1 . 1 . . . . . 
    . . . . . . . . . . . . . . . . 1 . 1 . 1 . . . 2 . . . . . . . . . 1 1 . . . . . 1 . . . 1 . . . . . . . . 
    . . . . . . . . . . . . . . . . . . 1 . 1 . 1 1 1 1 2 1 1 . . . . . . 1 1 1 1 . . . . . . 1 . . . . . . . . 
    . . . . . . . . . . . . . . . . . 1 . . 1 . 1 1 . . . 1 . . . . . . . 1 . . 1 . . . . . . 1 1 1 1 1 1 1 . . 
    . . . . . . . . . . . . . . . . 1 1 . . . . 1 1 . . . . . . . . . . . . . . . . . . . . . 1 1 . . . . . . . 
    . . . . . . . . . . . . . . . . . 1 . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . 1 1 . . . . . . 
    . . . . . . . . . . . . . . . . . 1 . . 1 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . 2 1 1 . . 1 . . 1 1 1 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . 1 . . . . 1 . . 1 . . . . 1 . . 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . 1 . . . 1 1 1 1 1 . . . . 2 . . . 1 . 1 . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . 1 1 . . . 1 1 . 1 . . . . 1 . . . 1 . 1 . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 1 2 1 1 1 . . 1 1 . . . 1 . . . 1 1 1 . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . 1 2 1 1 1 . 1 1 1 1 1 . . . . . . . . 1 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . 1 1 . 1 1 . 1 . . 1 . . 1 . . . . . . . . . 1 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . 1 1 1 1 . . 1 . . 1 . . 1 1 . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . 2 . . 1 . . 2 . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . 1 2 . . . . . . . . . . . . . 1 1 1 . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . 1 1 1 1 1 . . . . . . . . . 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . 1 1 1 1 1 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . 1 1 1 . 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . 1 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 
    . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . 

## Benchmarks
1. go test -bench=".*"

### Benchmark example results

BenchmarkS01           1        37960644383 ns/op

BenchmarkS02           1        33688144584 ns/op

BenchmarkS03           1        35811056919 ns/op

## Profileing with testing package
1. go test -bench=".*" -cpuprofile walk.cpu.out
2. go tool pprof walk walk.cpu.out
3. top10
4. top10 -cum
5. web
6. weblist 
