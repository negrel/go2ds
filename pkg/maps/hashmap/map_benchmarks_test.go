package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"testing"
)

func benchmarkMap_Put(b *testing.B, size int) {
	m := New()
	keys := make([]int, size)
	for j := 0; j < size; j++ {
		keys[j] = j
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			key := maps.Key(&keys[j])
			value := maps.Value(&struct{}{})

			m.Put(key, value)
		}
		b.StopTimer()
		m.Clear()
		b.StartTimer()
	}
}

func benchmarkMap_Get(b *testing.B, size int) {
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

func benchmarkMap_Delete(b *testing.B, size int) {
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

func BenchmarkMap_Put_100(b *testing.B) {
	benchmarkMap_Put(b, 100)
}

func BenchmarkMap_Put_1000(b *testing.B) {
	benchmarkMap_Put(b, 1000)
}

func BenchmarkMap_Put_10000(b *testing.B) {
	benchmarkMap_Put(b, 10000)
}

func BenchmarkMap_Put_100000(b *testing.B) {
	benchmarkMap_Put(b, 100000)
}

func BenchmarkMap_Get_100(b *testing.B) {
	benchmarkMap_Get(b, 100)
}

func BenchmarkMap_Get_1000(b *testing.B) {
	benchmarkMap_Get(b, 1000)
}

func BenchmarkMap_Get_10000(b *testing.B) {
	benchmarkMap_Get(b, 10000)
}

func BenchmarkMap_Get_100000(b *testing.B) {
	benchmarkMap_Get(b, 100000)
}

func BenchmarkMap_Delete_100(b *testing.B) {
	benchmarkMap_Delete(b, 100)
}

func BenchmarkMap_Delete_1000(b *testing.B) {
	benchmarkMap_Delete(b, 1000)
}

func BenchmarkMap_Delete_10000(b *testing.B) {
	benchmarkMap_Delete(b, 10000)
}

func BenchmarkMap_Delete_100000(b *testing.B) {
	benchmarkMap_Delete(b, 100000)
}
