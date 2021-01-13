package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"testing"
)

func benchmarkMapSet(b *testing.B, size int) {
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

			m.Set(key, value)
		}
		b.StopTimer()
		m.Clear()
		b.StartTimer()
	}
}

func benchmarkMapGet(b *testing.B, size int) {
	m := New()
	keys := make([]maps.Key, size)

	for i := 0; i < size; i++ {
		j := i
		key := maps.Key(&j)
		value := maps.Value(&struct{}{})

		keys[i] = key

		m.Set(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, key := range keys {
			_, _ = m.Get(key)
		}
	}
}

func benchmarkMapDelete(b *testing.B, size int) {
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

			m.Set(key, value)
		}

		b.StartTimer()
		for _, key := range keys {
			m.Delete(key)
		}
	}
}

func BenchmarkMap_Set_100(b *testing.B) {
	benchmarkMapSet(b, 100)
}

func BenchmarkMap_Set_1000(b *testing.B) {
	benchmarkMapSet(b, 1000)
}

func BenchmarkMap_Set_10000(b *testing.B) {
	benchmarkMapSet(b, 10000)
}

func BenchmarkMap_Set_100000(b *testing.B) {
	benchmarkMapSet(b, 100000)
}

func BenchmarkMap_Get_100(b *testing.B) {
	benchmarkMapGet(b, 100)
}

func BenchmarkMap_Get_1000(b *testing.B) {
	benchmarkMapGet(b, 1000)
}

func BenchmarkMap_Get_10000(b *testing.B) {
	benchmarkMapGet(b, 10000)
}

func BenchmarkMap_Get_100000(b *testing.B) {
	benchmarkMapGet(b, 100000)
}

func BenchmarkMap_Delete_100(b *testing.B) {
	benchmarkMapDelete(b, 100)
}

func BenchmarkMap_Delete_1000(b *testing.B) {
	benchmarkMapDelete(b, 1000)
}

func BenchmarkMap_Delete_10000(b *testing.B) {
	benchmarkMapDelete(b, 10000)
}

func BenchmarkMap_Delete_100000(b *testing.B) {
	benchmarkMapDelete(b, 100000)
}
