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
go test -bench=Gammazero -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGammazeroWordsPut-8             34    30807080 ns/op    14173312 B/op    430510 allocs/op
BenchmarkGammazeroWordsGet-8             82    14091831 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWordsWalk-8           835     1380902 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWeb2aPut-8             44    25430734 ns/op    12538598 B/op    352464 allocs/op
BenchmarkGammazeroWeb2aGet-8             97    11865033 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWeb2aWalk-8           904     1386653 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroUUIDsPut-8             21    50514997 ns/op    15961753 B/op    437455 allocs/op
BenchmarkGammazeroUUIDsGet-8           1414      841096 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroUUIDsWalk-8         15552       76027 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroHSKPut-8              700     1712411 ns/op      781824 B/op     21603 allocs/op
BenchmarkGammazeroHSKGet-8             1422      843977 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroHSKWalk-8           15753       77257 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroPutWithExisting-8 3877678       321.7 ns/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    18.314s

go test -bench=Dghubble -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkDghubbleWordsPut-8              26    44399682 ns/op    34087599 B/op    668808 allocs/op
BenchmarkDghubbleWordsGet-8             100    10723301 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWordsWalk-8             52    21354546 ns/op     2388220 B/op    233697 allocs/op
BenchmarkDghubbleWeb2aPut-8              15    75496998 ns/op    71419068 B/op   1258426 allocs/op
BenchmarkDghubbleWeb2aGet-8             100    11418684 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aWalk-8             28    42157352 ns/op     5852674 B/op    442357 allocs/op
BenchmarkDghubbleUUIDsPut-8               2   598843178 ns/op   585421824 B/op   9658527 allocs/op
BenchmarkDghubbleUUIDsGet-8            4461      266545 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsWalk-8           2434      488783 ns/op       42018 B/op      6724 allocs/op
BenchmarkDghubbleHSKPut-8              1114     1152263 ns/op      732379 B/op     15975 allocs/op
BenchmarkDghubbleHSKGet-8              4542      266464 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKWalk-8             2451      487336 ns/op       42023 B/op      6724 allocs/op
BenchmarkDghubblePutWithExisting-8  4984672       248.0 ns/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    19.322s

go test -bench=GoRadix -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGoRadixWordsPut-8               28    42375949 ns/op    17050771 B/op    550402 allocs/op
BenchmarkGoRadixWordsGet-8               66    15907578 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalk-8             787     1466352 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aPut-8               33    38427931 ns/op    15178835 B/op    462473 allocs/op
BenchmarkGoRadixWeb2aGet-8               78    13547516 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalk-8             784     1542564 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsPut-8               18    69380880 ns/op    19259576 B/op    574864 allocs/op
BenchmarkGoRadixUUIDsGet-8             1228      929895 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalk-8           14708       80184 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKPut-8                488     2578328 ns/op      940386 B/op     28209 allocs/op
BenchmarkGoRadixHSKGet-8               1255      921039 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalk-8             15001       79543 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixPutWithExisting-8   2906332       453.1 ns/op
PASS
ok      github.com/gammazero/radixtree-compare-benchmark    19.980s

go test -bench=Plar -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkPlarWordsPut-8                  26    40974807 ns/op    16381322 B/op    568741 allocs/op
BenchmarkPlarWordsGet-8                  80    13257882 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWordsWalk-8                486     2344110 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWeb2aPut-8                  40    31399233 ns/op    13023783 B/op    427283 allocs/op
BenchmarkPlarWeb2aGet-8                 100    10653022 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aWalk-8                630     1865313 ns/op          40 B/op         2 allocs/op
BenchmarkPlarUUIDsPut-8                  20    61307567 ns/op    18838776 B/op    547647 allocs/op
BenchmarkPlarUUIDsGet-8                1866      621580 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsWalk-8               9632      122398 ns/op          40 B/op         2 allocs/op
BenchmarkPlarHSKPut-8                   675     1889763 ns/op      788260 B/op     27521 allocs/op
BenchmarkPlarHSKGet-8                  1844      609338 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKWalk-8                10000      118957 ns/op          40 B/op         2 allocs/op
BenchmarkPlarPutWithExisting-8      2870120       456.7 ns/op
PASS
ok  	github.com/gammazero/radixtree-compare-benchmark	17.259s
```

Generally a radixtree is built once and searched many times.  That makes the `Get` and `Walk` operations the most important; how fast can an item be looked up by its key and how fast can the tree be iterated. Here are only those values from above.
```
BenchmarkGammazeroWordsGet-8             82    14091831 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWordsWalk-8           835     1380902 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWeb2aGet-8             97    11865033 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroWeb2aWalk-8           904     1386653 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroUUIDsGet-8           1414      841096 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroUUIDsWalk-8         15552       76027 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroHSKGet-8             1422      843977 ns/op           0 B/op         0 allocs/op
BenchmarkGammazeroHSKWalk-8           15753       77257 ns/op           0 B/op         0 allocs/op

BenchmarkDghubbleWordsGet-8             100    10723301 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWordsWalk-8             52    21354546 ns/op     2388220 B/op    233697 allocs/op
BenchmarkDghubbleWeb2aGet-8             100    11418684 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aWalk-8             28    42157352 ns/op     5852674 B/op    442357 allocs/op
BenchmarkDghubbleUUIDsGet-8            4461      266545 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsWalk-8           2434      488783 ns/op       42018 B/op      6724 allocs/op
BenchmarkDghubbleHSKGet-8              4542      266464 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKWalk-8             2451      487336 ns/op       42023 B/op      6724 allocs/op

BenchmarkGoRadixWordsGet-8               66    15907578 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalk-8             787     1466352 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aGet-8               78    13547516 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalk-8             784     1542564 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsGet-8             1228      929895 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalk-8           14708       80184 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKGet-8               1255      921039 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalk-8             15001       79543 ns/op           0 B/op         0 allocs/op

BenchmarkPlarWordsGet-8                  80    13257882 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWordsWalk-8                486     2344110 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWeb2aGet-8                 100    10653022 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aWalk-8                630     1865313 ns/op          40 B/op         2 allocs/op
BenchmarkPlarUUIDsGet-8                1866      621580 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsWalk-8               9632      122398 ns/op          40 B/op         2 allocs/op
BenchmarkPlarHSKGet-8                  1844      609338 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKWalk-8                10000      118957 ns/op          40 B/op         2 allocs/op
```
