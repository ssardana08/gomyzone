## How to run benchmarks

### Running benchmarks with memory stats
```shell
cd benchmark
go test -benchmem -run ^$ -bench "^(BenchmarkReturnStruct|BenchmarkReturnStructPointer)$"
```

### Running benchmarks with memory stats and a time period
```shell
cd benchmark
go test -benchmem -run ^$ -bench "^(BenchmarkReturnStruct|BenchmarkReturnStructPointer)$" -benchtime=20s
```