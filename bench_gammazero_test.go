package triebench

import (
	"bufio"
	"os"
	"testing"

	"github.com/gammazero/radixtree"
)

const (
	wordsPath = "/usr/share/dict/words"
	web2aPath = "/usr/share/dict/web2a"
	uuidsPath = "uuids.txt"
	hskPath   = "hsk_words.txt"
)

//
// Benchmarks
//
func BenchmarkGammazeroWordsPut(b *testing.B) {
	benchmarkPut(wordsPath, b)
}

func BenchmarkGammazeroWordsGet(b *testing.B) {
	benchmarkGet(wordsPath, b)
}

func BenchmarkGammazeroWordsWalk(b *testing.B) {
	benchmarkWalk(wordsPath, b)
}

func BenchmarkGammazeroWeb2aPut(b *testing.B) {
	benchmarkPut(web2aPath, b)
}

func BenchmarkGammazeroWeb2aGet(b *testing.B) {
	benchmarkGet(web2aPath, b)
}

func BenchmarkGammazeroWeb2aWalk(b *testing.B) {
	benchmarkWalk(web2aPath, b)
}

func BenchmarkGammazeroUUIDsPut(b *testing.B) {
	benchmarkPut(uuidsPath, b)
}

func BenchmarkGammazeroUUIDsGet(b *testing.B) {
	benchmarkGet(hskPath, b)
}

func BenchmarkGammazeroUUIDsWalk(b *testing.B) {
	benchmarkWalk(hskPath, b)
}

func BenchmarkGammazeroHSKPut(b *testing.B) {
	benchmarkPut(hskPath, b)
}

func BenchmarkGammazeroHSKGet(b *testing.B) {
	benchmarkGet(hskPath, b)
}

func BenchmarkGammazeroHSKWalk(b *testing.B) {
	benchmarkWalk(hskPath, b)
}

func benchmarkPut(filePath string, b *testing.B) {
	words := loadWords(filePath)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := new(radixtree.Runes)
		for _, w := range words {
			tree.Put(w, w)
		}
	}
}

func benchmarkGet(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := new(radixtree.Runes)
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			tree.Get(w)
		}
	}
}

func benchmarkWalk(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := new(radixtree.Runes)
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		tree.Walk("", func(key string, value interface{}) error {
			count++
			return nil
		})
	}
	if count != len(words) {
		panic("wrong count")
	}
}

func loadWords(wordsFile string) []string {
	f, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var word string
	var words []string

	// Scan through line-dilimited words.
	for scanner.Scan() {
		word = scanner.Text()
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}
