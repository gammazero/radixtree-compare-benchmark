# radixtree-compare-benchmark
Compare benchmarks of radixtree and trie implementations

## Run benchmarks
```
git clone git@github.com:ajgillis/radixtree-compare-benchmark.git
cd radixtree-compare-benchmark
make
```

## Conclusions
Radix tree implementations have better performance, in memory and time when compared with a trie, for data sets where many intermediate nodes can be compressed into prefix data.  If the data cannot be compressed (no common prefixes), then the trie and radix tree are equivalent, with a slight advantage to the trie for simplicity and because it does not create any compressed nodes when building the tree.  Both radix tree implementations, compared here, have nearly the same performance.

For reggie, the the gammazero radixtree implementation has the advantage that it provides a `WalkPath` API that allows retrieval of all values in the path from root to the location of a key.

## Benchmark Results
```
go test -v -bench=Dghubble -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkDghubbleWordsPut-12                   8         134543693 ns/op        125885938 B/op   2229142 allocs/op
BenchmarkDghubbleWordsGet-12                  38          28549657 ns/op               0 B/op          0 allocs/op
BenchmarkDghubbleWordsWalk-12                 15          77109279 ns/op         9161249 B/op     792724 allocs/op
BenchmarkDghubbleWeb2aPut-12                  16          69230016 ns/op        74957781 B/op    1258424 allocs/op
BenchmarkDghubbleWeb2aGet-12                 100          11028805 ns/op               0 B/op          0 allocs/op
BenchmarkDghubbleWeb2aWalk-12                 30          39482565 ns/op         5955901 B/op     442305 allocs/op
BenchmarkDghubbleUUIDsPut-12                   2         554695238 ns/op        611426016 B/op   9658519 allocs/op
BenchmarkDghubbleUUIDsGet-12                3847            290867 ns/op               0 B/op          0 allocs/op
BenchmarkDghubbleUUIDsWalk-12               2408            501537 ns/op           42020 B/op       6724 allocs/op
BenchmarkDghubbleHSKPut-12                  1086           1097552 ns/op          786126 B/op      15974 allocs/op
BenchmarkDghubbleHSKGet-12                  4090            295192 ns/op               0 B/op          0 allocs/op
BenchmarkDghubbleHSKWalk-12                 1953            614740 ns/op           80848 B/op       5876 allocs/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    17.360s

go test -v -bench=Gammazero -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkGammazeroWordsPut-12                 16          69873499 ns/op        42883280 B/op     998942 allocs/op
BenchmarkGammazeroWordsGet-12                 39          29112193 ns/op               0 B/op          0 allocs/op
BenchmarkGammazeroWordsWalk-12                27          46592921 ns/op         3663363 B/op     318742 allocs/op
BenchmarkGammazeroWeb2aPut-12                 48          23521703 ns/op        14726276 B/op     335883 allocs/op
BenchmarkGammazeroWeb2aGet-12                123           9652358 ns/op               0 B/op          0 allocs/op
BenchmarkGammazeroWeb2aWalk-12                74          14686684 ns/op         1522595 B/op     109957 allocs/op
BenchmarkGammazeroUUIDsPut-12                 22          52356775 ns/op        28749123 B/op     418816 allocs/op
BenchmarkGammazeroUUIDsGet-12               3626            310241 ns/op               0 B/op          0 allocs/op
BenchmarkGammazeroUUIDsWalk-12              1947            616268 ns/op           80848 B/op       5876 allocs/op
BenchmarkGammazeroHSKPut-12                 1005           1192040 ns/op          730673 B/op      17855 allocs/op
BenchmarkGammazeroHSKGet-12                 3654            311534 ns/op               0 B/op          0 allocs/op
BenchmarkGammazeroHSKWalk-12                1968            615005 ns/op           80848 B/op       5876 allocs/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    17.139s

go test -v -bench=Plar -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkPlarWordsPut-12              16          70071335 ns/op        42883639 B/op     998946 allocs/op
BenchmarkPlarWordsGet-12              38          29572952 ns/op               0 B/op          0 allocs/op
BenchmarkPlarWordsWalk-12             27          48006743 ns/op         3663700 B/op     318742 allocs/op
BenchmarkPlarWeb2aPut-12              48          23553890 ns/op        14726223 B/op     335882 allocs/op
BenchmarkPlarWeb2aGet-12             120           9756100 ns/op               0 B/op          0 allocs/op
BenchmarkPlarWeb2aWalk-12             73          14215845 ns/op         1522939 B/op     109957 allocs/op
BenchmarkPlarUUIDsPut-12              21          50913627 ns/op        28748304 B/op     418809 allocs/op
BenchmarkPlarUUIDsGet-12            3534            324518 ns/op               0 B/op          0 allocs/op
BenchmarkPlarUUIDsWalk-12           1971            613734 ns/op           80848 B/op       5876 allocs/op
BenchmarkPlarHSKPut-12               973           1192070 ns/op          730640 B/op      17854 allocs/op
BenchmarkPlarHSKGet-12              3783            313821 ns/op               0 B/op          0 allocs/op
BenchmarkPlarHSKWalk-12             1885            613273 ns/op           80848 B/op       5876 allocs/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    17.058s
```
