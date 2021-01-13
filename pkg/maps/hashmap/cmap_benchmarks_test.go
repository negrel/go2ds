package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"sync"
	"testing"
)

func benchmarkCMap_Put(b *testing.B, size int) {
	m := NewC()
	keys := make([]int, size)
	for j := 0; j < size; j++ {
		keys[j] = j
	}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(size)
		for j := 0; j < size; j++ {
			go func(j int) {
				defer wg.Done()

				key := maps.Key(&keys[j])
				value := maps.Value(&struct{}{})

				m.Put(key, value)
			}(j)
		}
		wg.Wait()

		b.StopTimer()
		m.Clear()
		b.StartTimer()
	}
}

func benchmarkCMap_Get(b *testing.B, size int) {
	m := NewC()
	keys := make([]maps.Key, size)

	for i := 0; i < size; i++ {
		j := i
		key := maps.Key(&j)
		value := maps.Value(&struct{}{})

		keys[i] = key

		m.Put(key, value)
	}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(size)
		for _, key := range keys {
			go func(key maps.Key) {
				defer wg.Done()

				_, _ = m.Get(key)
			}(key)
		}
	}
	wg.Wait()
}

func benchmarkCMap_Delete(b *testing.B, size int) {
	m := NewC()
	keys := make([]maps.Key, size)

	for j := 0; j < size; j++ {
		k := j
		keys[j] = maps.Key(&k)
	}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		for j := 0; j < size; j++ {
			value := maps.Value(&struct{}{})
			m.Put(keys[j], value)
		}
		b.StartTimer()

		wg.Add(size)
		for _, key := range keys {
			go func(key maps.Key) {
				defer wg.Done()
				m.Delete(key)
			}(key)
		}
		wg.Wait()
	}
}

func BenchmarkCMap_Put_100(b *testing.B) {
	benchmarkCMap_Put(b, 100)
}

func BenchmarkCMap_Put_1000(b *testing.B) {
	benchmarkCMap_Put(b, 1000)
}

func BenchmarkCMap_Put_10000(b *testing.B) {
	benchmarkCMap_Put(b, 10000)
}

func BenchmarkCMap_Put_100000(b *testing.B) {
	benchmarkCMap_Put(b, 100000)
}

func BenchmarkCMap_Get_100(b *testing.B) {
	benchmarkCMap_Get(b, 100)
}

func BenchmarkCMap_Get_1000(b *testing.B) {
	benchmarkCMap_Get(b, 1000)
}

func BenchmarkCMap_Get_10000(b *testing.B) {
	benchmarkCMap_Get(b, 10000)
}

func BenchmarkCMap_Get_100000(b *testing.B) {
	benchmarkCMap_Get(b, 100000)
}

func BenchmarkCMap_Delete_100(b *testing.B) {
	benchmarkCMap_Delete(b, 100)
}

func BenchmarkCMap_Delete_1000(b *testing.B) {
	benchmarkCMap_Delete(b, 1000)
}

func BenchmarkCMap_Delete_10000(b *testing.B) {
	benchmarkCMap_Delete(b, 10000)
}

func BenchmarkCMap_Delete_100000(b *testing.B) {
	benchmarkCMap_Delete(b, 100000)
}
