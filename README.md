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
go test -bench=Dghubble -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkDghubbleWordsPut-12     	       8	 132737790 ns/op	125884195 B/op	 2229127 allocs/op
BenchmarkDghubbleWordsGet-12     	      39	  30267061 ns/op	       0 B/op	       0 allocs/op
BenchmarkDghubbleWordsWalk-12    	      16	  79401823 ns/op	 9160496 B/op	  792724 allocs/op
BenchmarkDghubbleWeb2aPut-12     	      15	  69461664 ns/op	74958051 B/op	 1258427 allocs/op
BenchmarkDghubbleWeb2aGet-12     	      93	  13843002 ns/op	       0 B/op	       0 allocs/op
BenchmarkDghubbleWeb2aWalk-12    	      28	  40994048 ns/op	 5955848 B/op	  442305 allocs/op
BenchmarkDghubbleUUIDsPut-12     	       2	 544645704 ns/op	611424728 B/op	 9658507 allocs/op
BenchmarkDghubbleUUIDsGet-12     	    4279	    284080 ns/op	       0 B/op	       0 allocs/op
BenchmarkDghubbleUUIDsWalk-12    	    2404	    496382 ns/op	   42058 B/op	    6724 allocs/op
BenchmarkDghubbleHSKPut-12       	    1098	   1107079 ns/op	  786198 B/op	   15975 allocs/op
BenchmarkDghubbleHSKGet-12       	    4218	    286492 ns/op	       0 B/op	       0 allocs/op
BenchmarkDghubbleHSKWalk-12      	    5406	    219947 ns/op	      87 B/op	       3 allocs/op
PASS
ok  	github.com/gammazero/radixtree-compare-benchmark	17.626s

go test -bench=Gammazero -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkGammazeroWordsPut-12     	      15	  70197672 ns/op	42883810 B/op	  998948 allocs/op
BenchmarkGammazeroWordsGet-12     	      37	  30377617 ns/op	       0 B/op	       0 allocs/op
BenchmarkGammazeroWordsWalk-12    	      61	  22729088 ns/op	     268 B/op	       5 allocs/op
BenchmarkGammazeroWeb2aPut-12     	      43	  23392203 ns/op	14726405 B/op	  335884 allocs/op
BenchmarkGammazeroWeb2aGet-12     	     121	  10188664 ns/op	       0 B/op	       0 allocs/op
BenchmarkGammazeroWeb2aWalk-12    	     214	   5652795 ns/op	     245 B/op	       5 allocs/op
BenchmarkGammazeroUUIDsPut-12     	      24	  54424759 ns/op	28748823 B/op	  418814 allocs/op
BenchmarkGammazeroUUIDsGet-12     	    3762	    307418 ns/op	       0 B/op	       0 allocs/op
BenchmarkGammazeroUUIDsWalk-12    	    5660	    212842 ns/op	      87 B/op	       3 allocs/op
BenchmarkGammazeroHSKPut-12       	     987	   1203088 ns/op	  730620 B/op	   17854 allocs/op
BenchmarkGammazeroHSKGet-12       	    3835	    306555 ns/op	       0 B/op	       0 allocs/op
BenchmarkGammazeroHSKWalk-12      	    5758	    217624 ns/op	      87 B/op	       3 allocs/op
PASS
ok  	github.com/gammazero/radixtree-compare-benchmark	18.100s

go test -bench=Plar -run=xx
goos: darwin
goarch: amd64
pkg: github.com/gammazero/radixtree-compare-benchmark
BenchmarkPlarWordsPut-12     	      13	  85378830 ns/op	41299892 B/op	 1326167 allocs/op
BenchmarkPlarWordsGet-12     	      36	  32499140 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlarWordsWalk-12    	     165	   7283012 ns/op	      56 B/op	       3 allocs/op
BenchmarkPlarWeb2aPut-12     	      43	  27253088 ns/op	13675841 B/op	  427284 allocs/op
BenchmarkPlarWeb2aGet-12     	     100	  10355546 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlarWeb2aWalk-12    	     615	   1974335 ns/op	      56 B/op	       3 allocs/op
BenchmarkPlarUUIDsPut-12     	      19	  58843893 ns/op	19638823 B/op	  547648 allocs/op
BenchmarkPlarUUIDsGet-12     	    1977	    616156 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlarUUIDsWalk-12    	    9243	    131229 ns/op	      56 B/op	       3 allocs/op
BenchmarkPlarHSKPut-12       	     711	   1697263 ns/op	  828253 B/op	   27522 allocs/op
BenchmarkPlarHSKGet-12       	    1981	    594770 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlarHSKWalk-12      	    9632	    128923 ns/op	      56 B/op	       3 allocs/op
PASS
ok  	github.com/gammazero/radixtree-compare-benchmark	17.474s
```

Generally a radixtree is built once and searched many times.  That makes the `Get` operation the most important; how fast can an item be looked up by its key.  Here are only the `Get` operations from above:

```
BenchmarkDghubbleWordsGet-12       39   30267061 ns/op   0 B/op   0 allocs/op
BenchmarkDghubbleWeb2aGet-12       93   13843002 ns/op   0 B/op   0 allocs/op
BenchmarkDghubbleUUIDsGet-12     4279     284080 ns/op   0 B/op   0 allocs/op
BenchmarkDghubbleHSKGet-12       4218     286492 ns/op   0 B/op   0 allocs/op

BenchmarkGammazeroWordsGet-12      37   30377617 ns/op   0 B/op   0 allocs/op
BenchmarkGammazeroWeb2aGet-12     121   10188664 ns/op   0 B/op   0 allocs/op
BenchmarkGammazeroUUIDsGet-12    3762     307418 ns/op   0 B/op   0 allocs/op
BenchmarkGammazeroHSKGet-12      3835     306555 ns/op   0 B/op   0 allocs/op

BenchmarkPlarWordsGet-12           36   32499140 ns/op    0 B/op   0 allocs/op
BenchmarkPlarWeb2aGet-12          100   10355546 ns/op    0 B/op   0 allocs/op
BenchmarkPlarUUIDsGet-12         1977     616156 ns/op    0 B/op   0 allocs/op
BenchmarkPlarHSKGet-12           1981     594770 ns/op    0 B/op   0 allocs/op
```

The deeper the tree (more keys with a common prefix) the better the gammazero radixtree implementation does.  Since the geohashes we use are very hierarchical, the gammazero radixtree implementation is expected to consistently exhibit the best performance. This is the implementation that Reggie uses.
