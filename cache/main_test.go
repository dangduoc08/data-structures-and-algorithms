package cache

import (
	"testing"
)

func TestInsert(t *testing.T) {
	inputs := []int{
		35,
		33,
		42,
		10,
		14,
		19,
		27,
		44,
		26,
		31,
	}

	heap := []int{
		10,
		14,
		19,
		26,
		31,
		42,
		27,
		44,
		35,
		33,
	}

	minHeap := NewMinHeap(inputs[0])
	for i := 1; i < len(inputs); i++ {
		value := inputs[i]
		minHeap.Insert(value)
	}

	for i, v := range minHeap.list {
		if v != heap[i] {
			t.Errorf("Value from list = %v not matched with expected value = %v", v, heap[i])
		}
	}
}

func TestRemoveMint(t *testing.T) {
	inputs := []int{
		35,
		33,
		42,
		10,
		14,
		19,
		27,
		44,
		26,
		31,
	}
	minHeap := NewMinHeap(inputs[0])
	for i := 1; i < len(inputs); i++ {
		value := inputs[i]
		minHeap.Insert(value)
	}

	rmHeap1 := []int{
		14,
		26,
		19,
		33,
		31,
		42,
		27,
		44,
		35,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap1[i] {
			t.Errorf("Value from removed list 1 = %v not matched with expected value = %v", v, rmHeap1[i])
		}
	}

	rmHeap2 := []int{
		19,
		26,
		27,
		33,
		31,
		42,
		35,
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap2[i] {
			t.Errorf("Value from removed list 2 = %v not matched with expected value = %v", v, rmHeap2[i])
		}
	}

	rmHeap3 := []int{
		26,
		31,
		27,
		33,
		44,
		42,
		35,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap3[i] {
			t.Errorf("Value from removed list 3 = %v not matched with expected value = %v", v, rmHeap3[i])
		}
	}

	rmHeap4 := []int{
		27,
		31,
		35,
		33,
		44,
		42,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap4[i] {
			t.Errorf("Value from removed list 4 = %v not matched with expected value = %v", v, rmHeap4[i])
		}
	}

	rmHeap5 := []int{
		31,
		33,
		35,
		42,
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap5[i] {
			t.Errorf("Value from removed list 5 = %v not matched with expected value = %v", v, rmHeap5[i])
		}
	}

	rmHeap6 := []int{
		33,
		42,
		35,
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap6[i] {
			t.Errorf("Value from removed list 6 = %v not matched with expected value = %v", v, rmHeap6[i])
		}
	}

	rmHeap7 := []int{
		35,
		42,
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap7[i] {
			t.Errorf("Value from removed list 7 = %v not matched with expected value = %v", v, rmHeap7[i])
		}
	}

	rmHeap8 := []int{
		42,
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap8[i] {
			t.Errorf("Value from removed list 8 = %v not matched with expected value = %v", v, rmHeap8[i])
		}
	}

	rmHeap9 := []int{
		44,
	}
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap9[i] {
			t.Errorf("Value from removed list 9 = %v not matched with expected value = %v", v, rmHeap9[i])
		}
	}

	rmHeap10 := []int{}
	minHeap.RemoveMin()
	minHeap.RemoveMin()
	for i, v := range minHeap.list {
		if v != rmHeap10[i] {
			t.Errorf("Value from removed list 10 = %v not matched with expected value = %v", v, rmHeap10[i])
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	minHeap := NewMinHeap(268136789173891)

	for i := 0; i < b.N; i++ {
		minHeap.Insert(i)
	}

	// BenchmarkInsert-12      84739562                23.88 ns/op           43 B/op          0 allocs/op
}

func BenchmarkRemoveMin(b *testing.B) {
	minHeap := NewMinHeap(1000000)

	for i := 1000000; i >= 0; i-- {
		minHeap.Insert(i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		minHeap.RemoveMin()
	}

	// BenchmarkRemoveMin-12           687878089                1.712 ns/op           0 B/op          0 allocs/op
}
