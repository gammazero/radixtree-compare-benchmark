package triebench

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	gart "github.com/plar/go-adaptive-radix-tree"
)

//
// Benchmarks
//
func BenchmarkPlarWordsPut(b *testing.B) {
	benchmarkPlarPut(wordsPath, b)
}

func BenchmarkPlarWordsGet(b *testing.B) {
	benchmarkPlarGet(wordsPath, b)
}

func BenchmarkPlarWordsWalk(b *testing.B) {
	benchmarkPlarWalk(wordsPath, b)
}

func BenchmarkPlarWordsWalkPath(b *testing.B) {
	benchmarkPlarWalkPath(wordsPath, b)
}

func BenchmarkPlarWeb2aPut(b *testing.B) {
	benchmarkPlarPut(web2aPath, b)
}

func BenchmarkPlarWeb2aGet(b *testing.B) {
	benchmarkPlarGet(web2aPath, b)
}

func BenchmarkPlarWeb2aWalk(b *testing.B) {
	benchmarkPlarWalk(web2aPath, b)
}

func BenchmarkPlarWeb2aWalkPath(b *testing.B) {
	benchmarkPlarWalkPath(web2aPath, b)
}

func BenchmarkPlarUUIDsPut(b *testing.B) {
	benchmarkPlarPut(uuidsPath, b)
}

func BenchmarkPlarUUIDsGet(b *testing.B) {
	benchmarkPlarGet(uuidsPath, b)
}

func BenchmarkPlarUUIDsWalk(b *testing.B) {
	benchmarkPlarWalk(uuidsPath, b)
}

func BenchmarkPlarUUIDsWalkPath(b *testing.B) {
	benchmarkPlarWalkPath(uuidsPath, b)
}

func BenchmarkPlarHSKPut(b *testing.B) {
	benchmarkPlarPut(hskPath, b)
}

func BenchmarkPlarHSKGet(b *testing.B) {
	benchmarkPlarGet(hskPath, b)
}

func BenchmarkPlarHSKWalk(b *testing.B) {
	benchmarkPlarWalk(hskPath, b)
}

func BenchmarkPlarHSKWalkPath(b *testing.B) {
	benchmarkPlarWalkPath(hskPath, b)
}

func BenchmarkPlarPutWithExisting(b *testing.B) {
	tree := gart.New()
	for i := 0; i < 10000; i++ {
		tree.Insert([]byte(fmt.Sprintf("somekey%d", i)), true)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		_, updated := tree.Insert([]byte(strconv.Itoa(n)), true)
		if updated {
			b.Fatal("value was overwritten")
		}
	}
}

func benchmarkPlarPut(filePath string, b *testing.B) {
	words := loadWordsBytes(filePath)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := gart.New()
		for _, w := range words {
			tree.Insert(w, w)
		}
	}
}

func benchmarkPlarGet(filePath string, b *testing.B) {
	words := loadWordsBytes(filePath)
	tree := gart.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			tree.Search(w)
		}
	}
}

func benchmarkPlarWalk(filePath string, b *testing.B) {
	words := loadWordsBytes(filePath)
	tree := gart.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		tree.ForEach(func(n gart.Node) bool {
			count++
			return true
		}, gart.TraverseLeaf)
	}
	if count != len(words) {
		panic("wrong count")
	}
}

func benchmarkPlarWalkPath(filePath string, b *testing.B) {
	words := loadWordsBytes(filePath)
	tree := gart.New()
	for _, w := range words {
		tree.Insert(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		for _, w := range words {
			tree.ForEachPrefix(w, func(n gart.Node) bool {
				count++
				return true
			})
		}
	}
	if count < len(words) {
		panic("wrong count")
	}
}

func loadWordsBytes(wordsFile string) [][]byte {
	f, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var word string
	var words [][]byte

	// Scan through line-dilimited words.
	for scanner.Scan() {
		word = scanner.Text()
		words = append(words, []byte(word))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}
