package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"testing"
)

func benchmarkPut(b *testing.B, size int) {
	m := New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			k := j
			key := maps.Key(&k)
			value := maps.Value(&struct{}{})

			m.Put(key, value)
		}
	}
}

func benchmarkGet(b *testing.B, size int) {
	m := New()
	keys := make([]maps.Key, size)

	for i := 0; i < size; i++ {
		j := i
		key := maps.Key(&j)
		value := maps.Value(&struct{}{})

		keys[i] = key

		m.Put(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			_, _ = m.Get(key)
		}
	}
}

func benchmarkDelete(b *testing.B, size int) {
	m := New()
	keys := make([]maps.Key, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		for j := 0; j < size; j++ {
			k := j
			key := maps.Key(&k)
			value := maps.Value(&struct{}{})

			keys[j] = key

			m.Put(key, value)
		}

		b.StartTimer()
		for _, key := range keys {
			m.Delete(key)
		}
	}
}

func BenchmarkHashMap_Put_100(b *testing.B) {
	benchmarkPut(b, 100)
}

func BenchmarkHashMap_Put_1000(b *testing.B) {
	benchmarkPut(b, 1000)
}

func BenchmarkHashMap_Put_10000(b *testing.B) {
	benchmarkPut(b, 10000)
}

func BenchmarkHashMap_Put_100000(b *testing.B) {
	benchmarkPut(b, 100000)
}

func BenchmarkHashMap_Get_100(b *testing.B) {
	benchmarkGet(b, 100)
}

func BenchmarkHashMap_Get_1000(b *testing.B) {
	benchmarkGet(b, 1000)
}

func BenchmarkHashMap_Get_10000(b *testing.B) {
	benchmarkGet(b, 10000)
}

func BenchmarkHashMap_Get_100000(b *testing.B) {
	benchmarkGet(b, 100000)
}

func BenchmarkHashMap_Delete_100(b *testing.B) {
	benchmarkDelete(b, 100)
}

func BenchmarkHashMap_Delete_1000(b *testing.B) {
	benchmarkDelete(b, 1000)
}

func BenchmarkHashMap_Delete_10000(b *testing.B) {
	benchmarkDelete(b, 10000)
}

func BenchmarkHashMap_Delete_100000(b *testing.B) {
	benchmarkDelete(b, 100000)
}
