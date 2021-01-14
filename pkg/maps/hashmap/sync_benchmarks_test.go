package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"sync"
	"testing"
)

func benchmarkSyncPut(b *testing.B, size int) {
	m := NewSync()
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

				m.Set(key, value)
			}(j)
		}
		wg.Wait()

		b.StopTimer()
		m.Clear()
		b.StartTimer()
	}
}

func benchmarkSyncGet(b *testing.B, size int) {
	m := NewSync()
	keys := make([]maps.Key, size)

	for i := 0; i < size; i++ {
		j := i
		key := maps.Key(&j)
		value := maps.Value(&struct{}{})

		keys[i] = key

		m.Set(key, value)
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

func benchmarkSyncDelete(b *testing.B, size int) {
	m := NewSync()
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
			m.Set(keys[j], value)
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

func BenchmarkSync_Put_100(b *testing.B) {
	benchmarkSyncPut(b, 100)
}

func BenchmarkSync_Put_1000(b *testing.B) {
	benchmarkSyncPut(b, 1000)
}

func BenchmarkSync_Put_10000(b *testing.B) {
	benchmarkSyncPut(b, 10000)
}

func BenchmarkSync_Put_100000(b *testing.B) {
	benchmarkSyncPut(b, 100000)
}

func BenchmarkSync_Get_100(b *testing.B) {
	benchmarkSyncGet(b, 100)
}

func BenchmarkSync_Get_1000(b *testing.B) {
	benchmarkSyncGet(b, 1000)
}

func BenchmarkSync_Get_10000(b *testing.B) {
	benchmarkSyncGet(b, 10000)
}

func BenchmarkSync_Get_100000(b *testing.B) {
	benchmarkSyncGet(b, 100000)
}

func BenchmarkSync_Delete_100(b *testing.B) {
	benchmarkSyncDelete(b, 100)
}

func BenchmarkSync_Delete_1000(b *testing.B) {
	benchmarkSyncDelete(b, 1000)
}

func BenchmarkSync_Delete_10000(b *testing.B) {
	benchmarkSyncDelete(b, 10000)
}

func BenchmarkSync_Delete_100000(b *testing.B) {
	benchmarkSyncDelete(b, 100000)
}
