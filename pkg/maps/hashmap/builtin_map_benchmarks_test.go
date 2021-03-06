package hashmap

import (
	"testing"
)

func benchmarkBuiltInSet(b *testing.B, size int) {
	m := make(map[int]struct{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			m[j] = struct{}{}
		}
		b.StopTimer()
		m = make(map[int]struct{})
		b.StartTimer()
	}
}

func benchmarkBuiltInGet(b *testing.B, size int) {
	m := make(map[int]struct{})

	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			m[j] = struct{}{}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			_, _ = m[j]
		}
	}
}

func benchmarkBuiltInDelete(b *testing.B, size int) {
	m := make(map[int]struct{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		for j := 0; j < size; j++ {
			m[j] = struct{}{}
		}

		b.StartTimer()
		for j := 0; j < size; j++ {
			delete(m, j)
		}
	}
}

func BenchmarkBuiltInMap_Set_100(b *testing.B) {
	benchmarkBuiltInSet(b, 100)
}

func BenchmarkBuiltInMap_Set_1000(b *testing.B) {
	benchmarkBuiltInSet(b, 1000)
}

func BenchmarkBuiltInMap_Set_10000(b *testing.B) {
	benchmarkBuiltInSet(b, 10000)
}

func BenchmarkBuiltInMap_Set_100000(b *testing.B) {
	benchmarkBuiltInSet(b, 100000)
}

func BenchmarkBuiltInMap_Get_100(b *testing.B) {
	benchmarkBuiltInGet(b, 100)
}

func BenchmarkBuiltInMap_Get_1000(b *testing.B) {
	benchmarkBuiltInGet(b, 1000)
}

func BenchmarkBuiltInMap_Get_10000(b *testing.B) {
	benchmarkBuiltInGet(b, 10000)
}

func BenchmarkBuiltInMap_Get_100000(b *testing.B) {
	benchmarkBuiltInGet(b, 100000)
}

func BenchmarkBuiltInMap_Delete_100(b *testing.B) {
	benchmarkBuiltInDelete(b, 100)
}

func BenchmarkBuiltInMap_Delete_1000(b *testing.B) {
	benchmarkBuiltInDelete(b, 1000)
}

func BenchmarkBuiltInMap_Delete_10000(b *testing.B) {
	benchmarkBuiltInDelete(b, 10000)
}

func BenchmarkBuiltInMap_Delete_100000(b *testing.B) {
	benchmarkBuiltInDelete(b, 100000)
}
