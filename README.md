# radixtree-compare-benchmark
Compare benchmarks of radixtree and trie implementations

## Run benchmarks
```
git clone git@github.com:gammazero/radixtree-compare-benchmark.git
cd radixtree-compare-benchmark
make
```

## Conclusions
Radix tree implementations have better performance, in memory and time when compared with a trie, for data sets where many intermediate nodes can be compressed into prefix data.  If the data cannot be compressed (no common prefixes), then the trie and radix tree are equivalent, with a slight advantage to the trie for simplicity and because it does not create any compressed nodes when building the tree.

Of these implementations, only the gammazero and goradix implementations have zero allocations for all read APIs.

## Benchmark Results

Implementation: https://github.com/gammazero/radixtree v0.2.3
```
go test -bench=Gammazero -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGammazeroWordsPut-8             45    25693760 ns/op    14173258 B/op   430510 allocs/op
BenchmarkGammazeroWordsGet-8            145     8207558 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWordsWalk-8           866     1342711 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWordsWalkPath-8       132     8979272 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aPut-8             55    21540917 ns/op    12538611 B/op   352464 allocs/op
BenchmarkGammazeroWeb2aGet-8            172     7038219 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aWalk-8           916     1530802 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aWalkPath-8       164     7202914 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsPut-8             24    44664276 ns/op    15961780 B/op   437455 allocs/op
BenchmarkGammazeroUUIDsGet-8           2233      518530 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsWalk-8         15111       76827 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsWalkPath-8        45    25556481 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKPut-8              842     1449894 ns/op      781824 B/op    21603 allocs/op
BenchmarkGammazeroHSKGet-8             2320      509658 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKWalk-8           15790       76853 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKWalkPath-8        2307      526657 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroPutWithExisting-8 5164003       242.3 ns/op         137 B/op        3 allocs/op
```

Implementation: https://github.com/armon/go-radix v1.0.0
```
go test -bench=GoRadix -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGoRadixWordsPut-8               28    42375949 ns/op    17050771 B/op    550402 allocs/op
BenchmarkGoRadixWordsGet-8               66    15907578 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalk-8             787     1466352 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalkPath-8          76    16087867 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aPut-8               33    38427931 ns/op    15178835 B/op    462473 allocs/op
BenchmarkGoRadixWeb2aGet-8               78    13547516 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalk-8             784     1542564 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalkPath-8          92    13049952 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsPut-8               18    69380880 ns/op    19259576 B/op    574864 allocs/op
BenchmarkGoRadixUUIDsGet-8             1228      929895 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalk-8           14708       80184 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalkPath-8          31    36517267 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKPut-8                488     2578328 ns/op      940386 B/op     28209 allocs/op
BenchmarkGoRadixHSKGet-8               1255      921039 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalk-8             15001       79543 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalkPath-8          1287      952200 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixPutWithExisting-8   2914098       450.1 ns/op         161 B/op         4 allocs/op
```

Implementation: https://github.com/dghubble/trie v0.0.0-20210609182954-9a58e577d803
```
go test -bench=Dghubble -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkDghubbleWordsPut-8              26    44265471 ns/op    34087820 B/op    668809 allocs/op
BenchmarkDghubbleWordsGet-8             100    10865470 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWordsWalk-8             54    21330822 ns/op     2388184 B/op    233697 allocs/op
BenchmarkDghubbleWordsWalkPath-8         91    12453992 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aPut-8              14    75036416 ns/op    71419149 B/op   1258426 allocs/op
BenchmarkDghubbleWeb2aGet-8             100    11510034 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aWalk-8             27    41392222 ns/op     5852586 B/op    442357 allocs/op
BenchmarkDghubbleWeb2aWalkPath-8         86    12976935 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsPut-8               2   593092064 ns/op   585421824 B/op   9658527 allocs/op
BenchmarkDghubbleUUIDsGet-8              13    89637727 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsWalk-8              3   393708701 ns/op    79925456 B/op   3250643 allocs/op
BenchmarkDghubbleUUIDsWalkPath-8         12    95190883 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKPut-8              1095     1171249 ns/op      732367 B/op     15975 allocs/op
BenchmarkDghubbleHSKGet-8              4311      267976 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKWalk-8             2433      490514 ns/op       42103 B/op      6724 allocs/op
BenchmarkDghubbleHSKWalkPath-8         3812      306578 ns/op           0 B/op         0 allocs/op
BenchmarkDghubblePutWithExisting-8  4839458       258.6 ns/op          70 B/op         2 allocs/op
```

Implementation: https://github.com/plar/go-adaptive-radix-tree v1.0.4
```
go test -bench=Plar -run=xx
goos: linux
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkPlarWordsPut-8                  27    38648884 ns/op    15003037 B/op    519092 allocs/op
BenchmarkPlarWordsGet-8                  92    12217374 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWordsWalk-8                510     2259355 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWordsWalkPath-8             69    16961941 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aPut-8                  43    29529492 ns/op    11903951 B/op    387564 allocs/op
BenchmarkPlarWeb2aGet-8                 123     9602343 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aWalk-8                660     1881728 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWeb2aWalkPath-8            120    10627817 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsPut-8                  26    47339884 ns/op    17575683 B/op    485056 allocs/op
BenchmarkPlarUUIDsGet-8                  54    22727794 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsWalk-8                493     2390030 ns/op          40 B/op         2 allocs/op
BenchmarkPlarUUIDsWalkPath-8             55    21811770 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKPut-8                   847     1403741 ns/op      715581 B/op     24502 allocs/op
BenchmarkPlarHSKGet-8                  2779      466212 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKWalk-8                10000      112470 ns/op          40 B/op         2 allocs/op
BenchmarkPlarHSKWalkPath-8             2509      492133 ns/op           0 B/op         0 allocs/op
BenchmarkPlarPutWithExisting-8      3267259       406.5 ns/op         115 B/op         5 allocs/op
```

Generally a radixtree is built once and searched many times.  That makes the `Get` and `Walk` operations the most important; how fast can an item be looked up by its key and how fast can the tree be iterated. Here are only those values from above.
```
BenchmarkGammazeroWordsGet-8            145     8207558 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aGet-8            172     7038219 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsGet-8           2233      518530 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKGet-8             2320      509658 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWordsWalk-8           866     1342711 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aWalk-8           916     1530802 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsWalk-8         15111       76827 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKWalk-8           15790       76853 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWordsWalkPath-8       132     8979272 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroWeb2aWalkPath-8       164     7202914 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroUUIDsWalkPath-8        45    25556481 ns/op           0 B/op        0 allocs/op
BenchmarkGammazeroHSKWalkPath-8        2307      526657 ns/op           0 B/op        0 allocs/op

BenchmarkGoRadixWordsGet-8               66    15907578 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aGet-8               78    13547516 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsGet-8             1228      929895 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKGet-8               1255      921039 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalk-8             787     1466352 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalk-8             784     1542564 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalk-8           14708       80184 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalk-8             15001       79543 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWordsWalkPath-8          76    16087867 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixWeb2aWalkPath-8          92    13049952 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixUUIDsWalkPath-8          31    36517267 ns/op           0 B/op         0 allocs/op
BenchmarkGoRadixHSKWalkPath-8          1287      952200 ns/op           0 B/op         0 allocs/op

BenchmarkDghubbleWordsGet-8             100    10865470 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aGet-8             100    11510034 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsGet-8              13    89637727 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKGet-8              4311      267976 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWordsWalk-8             54    21330822 ns/op     2388184 B/op    233697 allocs/op
BenchmarkDghubbleWeb2aWalk-8             27    41392222 ns/op     5852586 B/op    442357 allocs/op
BenchmarkDghubbleUUIDsWalk-8              3   393708701 ns/op    79925456 B/op   3250643 allocs/op
BenchmarkDghubbleHSKWalk-8             2433      490514 ns/op       42103 B/op      6724 allocs/op
BenchmarkDghubbleWordsWalkPath-8         91    12453992 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleWeb2aWalkPath-8         86    12976935 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleUUIDsWalkPath-8         12    95190883 ns/op           0 B/op         0 allocs/op
BenchmarkDghubbleHSKWalkPath-8         3812      306578 ns/op           0 B/op         0 allocs/op

BenchmarkPlarWordsGet-8                  92    12217374 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aGet-8                 123     9602343 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsGet-8                  54    22727794 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKGet-8                  2779      466212 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWordsWalk-8                510     2259355 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWeb2aWalk-8                660     1881728 ns/op          40 B/op         2 allocs/op
BenchmarkPlarUUIDsWalk-8                493     2390030 ns/op          40 B/op         2 allocs/op
BenchmarkPlarHSKWalk-8                10000      112470 ns/op          40 B/op         2 allocs/op
BenchmarkPlarWordsWalkPath-8             69    16961941 ns/op           0 B/op         0 allocs/op
BenchmarkPlarWeb2aWalkPath-8            120    10627817 ns/op           0 B/op         0 allocs/op
BenchmarkPlarUUIDsWalkPath-8             55    21811770 ns/op           0 B/op         0 allocs/op
BenchmarkPlarHSKWalkPath-8             2509      492133 ns/op           0 B/op         0 allocs/op
```
