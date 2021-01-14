package hashmap

import (
	"sync"
	"testing"
)

func benchmarkSyncMapSet(b *testing.B, size int) {
	m := sync.Map{}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(size)
		for j := 0; j < size; j++ {
			go func(j int) {
				defer wg.Done()

				m.Store(j, struct{}{})
			}(j)
		}
		wg.Wait()

		b.StopTimer()
		m = sync.Map{}
		b.StartTimer()
	}
}

func benchmarkSyncMapGet(b *testing.B, size int) {
	m := sync.Map{}

	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			m.Store(j, struct{}{})
		}
	}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(size)

		for j := 0; j < size; j++ {
			go func(j int) {
				defer wg.Done()
				_, _ = m.Load(j)
			}(j)
		}
		wg.Wait()
	}
}

func benchmarkSyncMapDelete(b *testing.B, size int) {
	m := sync.Map{}

	var wg sync.WaitGroup

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		for j := 0; j < size; j++ {
			m.Store(j, struct{}{})
		}

		b.StartTimer()
		wg.Add(size)
		for j := 0; j < size; j++ {
			go func(j int) {
				defer wg.Done()

				m.Delete(j)
			}(j)
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap_Set_100(b *testing.B) {
	benchmarkSyncMapSet(b, 100)
}

func BenchmarkSyncMap_Set_1000(b *testing.B) {
	benchmarkSyncMapSet(b, 1000)
}

func BenchmarkSyncMap_Set_10000(b *testing.B) {
	benchmarkSyncMapSet(b, 10000)
}

func BenchmarkSyncMap_Set_100000(b *testing.B) {
	benchmarkSyncMapSet(b, 100000)
}

func BenchmarkSyncMap_Get_100(b *testing.B) {
	benchmarkSyncMapGet(b, 100)
}

func BenchmarkSyncMap_Get_1000(b *testing.B) {
	benchmarkSyncMapGet(b, 1000)
}

func BenchmarkSyncMap_Get_10000(b *testing.B) {
	benchmarkSyncMapGet(b, 10000)
}

func BenchmarkSyncMap_Get_100000(b *testing.B) {
	benchmarkSyncMapGet(b, 100000)
}

func BenchmarkSyncMap_Delete_100(b *testing.B) {
	benchmarkSyncMapDelete(b, 100)
}

func BenchmarkSyncMap_Delete_1000(b *testing.B) {
	benchmarkSyncMapDelete(b, 1000)
}

func BenchmarkSyncMap_Delete_10000(b *testing.B) {
	benchmarkSyncMapDelete(b, 10000)
}

func BenchmarkSyncMap_Delete_100000(b *testing.B) {
	benchmarkSyncMapDelete(b, 100000)
}
