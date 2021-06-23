package triebench

import (
	"fmt"
	"strconv"
	"testing"

	goradix "github.com/armon/go-radix"
)

//
// Benchmarks
//
func BenchmarkGoRadixWordsPut(b *testing.B) {
	benchmarkGoRadixPut(wordsPath, b)
}

func BenchmarkGoRadixWordsGet(b *testing.B) {
	benchmarkGoRadixGet(wordsPath, b)
}

func BenchmarkGoRadixWordsWalk(b *testing.B) {
	benchmarkGoRadixWalk(wordsPath, b)
}

func BenchmarkGoRadixWordsWalkPath(b *testing.B) {
	benchmarkGoRadixWalkPath(wordsPath, b)
}

func BenchmarkGoRadixWeb2aPut(b *testing.B) {
	benchmarkGoRadixPut(web2aPath, b)
}

func BenchmarkGoRadixWeb2aGet(b *testing.B) {
	benchmarkGoRadixGet(web2aPath, b)
}

func BenchmarkGoRadixWeb2aWalk(b *testing.B) {
	benchmarkGoRadixWalk(web2aPath, b)
}

func BenchmarkGoRadixWeb2aWalkPath(b *testing.B) {
	benchmarkGoRadixWalkPath(web2aPath, b)
}

func BenchmarkGoRadixUUIDsPut(b *testing.B) {
	benchmarkGoRadixPut(uuidsPath, b)
}

func BenchmarkGoRadixUUIDsGet(b *testing.B) {
	benchmarkGoRadixGet(uuidsPath, b)
}

func BenchmarkGoRadixUUIDsWalk(b *testing.B) {
	benchmarkGoRadixWalk(uuidsPath, b)
}

func BenchmarkGoRadixUUIDsWalkPath(b *testing.B) {
	benchmarkGoRadixWalkPath(uuidsPath, b)
}

func BenchmarkGoRadixHSKPut(b *testing.B) {
	benchmarkGoRadixPut(hskPath, b)
}

func BenchmarkGoRadixHSKGet(b *testing.B) {
	benchmarkGoRadixGet(hskPath, b)
}

func BenchmarkGoRadixHSKWalk(b *testing.B) {
	benchmarkGoRadixWalk(hskPath, b)
}

func BenchmarkGoRadixHSKWalkPath(b *testing.B) {
	benchmarkGoRadixWalkPath(hskPath, b)
}

func BenchmarkGoRadixPutWithExisting(b *testing.B) {
	tree := goradix.New()
	for i := 0; i < 10000; i++ {
		tree.Insert(fmt.Sprintf("somekey%d", i), true)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, updated := tree.Insert(strconv.Itoa(n), true)
		if updated {
			b.Fatal("value was overwritten")
		}
	}
}

func benchmarkGoRadixPut(filePath string, b *testing.B) {
	words := loadWords(filePath)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := goradix.New()
		for _, w := range words {
			tree.Insert(w, w)
		}
	}
}

func benchmarkGoRadixGet(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := goradix.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			tree.Get(w)
		}
	}
}

func benchmarkGoRadixWalk(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := goradix.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		tree.Walk(func(key string, value interface{}) bool {
			count++
			return false
		})
	}
	if count != len(words) {
		panic("wrong count")
	}
}

func benchmarkGoRadixWalkPath(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := goradix.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		for _, w := range words {
			tree.WalkPath(w, func(key string, value interface{}) bool {
				count++
				return false
			})
		}
	}
	if count < len(words) {
		panic("wrong count")
	}
}
