package triebench

import (
	"bufio"
	"os"
	"testing"

	gart "github.com/plar/go-adaptive-radix-tree"
)

//
// Benchmarks
//
func BenchmarkPlarWordsPut(b *testing.B) {
	benchmarkPut(wordsPath, b)
}

func BenchmarkPlarWordsGet(b *testing.B) {
	benchmarkGet(wordsPath, b)
}

func BenchmarkPlarWordsWalk(b *testing.B) {
	benchmarkWalk(wordsPath, b)
}

func BenchmarkPlarWeb2aPut(b *testing.B) {
	benchmarkPut(web2aPath, b)
}

func BenchmarkPlarWeb2aGet(b *testing.B) {
	benchmarkGet(web2aPath, b)
}

func BenchmarkPlarWeb2aWalk(b *testing.B) {
	benchmarkWalk(web2aPath, b)
}

func BenchmarkPlarUUIDsPut(b *testing.B) {
	benchmarkPut(uuidsPath, b)
}

func BenchmarkPlarUUIDsGet(b *testing.B) {
	benchmarkGet(hskPath, b)
}

func BenchmarkPlarUUIDsWalk(b *testing.B) {
	benchmarkWalk(hskPath, b)
}

func BenchmarkPlarHSKPut(b *testing.B) {
	benchmarkPut(hskPath, b)
}

func BenchmarkPlarHSKGet(b *testing.B) {
	benchmarkGet(hskPath, b)
}

func BenchmarkPlarHSKWalk(b *testing.B) {
	benchmarkWalk(hskPath, b)
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
