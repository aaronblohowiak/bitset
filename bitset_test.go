package bitset

import (
	"math/rand"
	"testing"
)

func TestNewBitset(t *testing.T) {
	var bits BitSet
	allocs := testing.AllocsPerRun(100, func() {
		bits = New(256)
	})

	if allocs != 2 {
		t.Error("Should pre-allocate bitset and backing array", allocs)
	}
}

func TestSetBit(t *testing.T) {
	b := New(100)

	allocs := testing.AllocsPerRun(100, func() {
		b.Set(0)
	})

	if allocs != 0 {
		t.Error("Should NOT allocate", allocs)
	}

	if b.Bytes()[0] != 1 {
		t.Error("After setting the 1st bit, the first byte should be 1, ", b.Bytes()[0])
	}

	b = New(32)
	b.Set(9)
	if b.Bytes()[1] != 2 {
		t.Error("After setting the 10th bit, the second byte should be 2, ", b.Bytes())
	}

	b = New(100)
	b.Set(64)
	if b.Bytes()[8] != 1 {
		t.Error("After setting the 65th bit, the second byte should be 1, ", b.Bytes())
	}
}

func TestGetBit(t *testing.T) {
	b := New(100)

	b.Set(65)

	if b.Get(0) != 0 {
		t.Error("An unset bit should be 0", b.Get(0))
	}

	if b.Get(65) != 1 {
		t.Error("A set bit should be 1", b.Bytes())
	}

	allocs := testing.AllocsPerRun(100, func() {
		b.Get(65)
	})

	if allocs != 0 {
		t.Error("Should NOT allocate", allocs)
	}
}

func TestClearBit(t *testing.T) {
	b := New(200)

	b.Set(65)

	if b.Get(65) != 1 {
		t.Error("A set bit should be 1", b.Bytes())
	}

	if b.Get(13) != 0 {
		t.Error("An unset bit should be 0", b.Bytes())
	}

	b.Clear(13)
	if b.Get(13) != 0 {
		t.Error("An unset bit should be 0", b.Bytes())
	}

	allocs := testing.AllocsPerRun(100, func() {
		b.Clear(65)
	})

	if b.Bytes()[8] != 0 {
		t.Error("A set bit should be 1", b)
	}

	if allocs != 0 {
		t.Error("Should NOT allocate", allocs)
	}

	if b.Get(65) != 0 {
		t.Error("A cleared bit should be 0", b)
	}

}

func TestAnd(t *testing.T) {
	a := New(100)
	a.Set(1)
	a.Set(27)

	b := New(100)
	b.Set(25)
	b.Set(27)

	allocs := testing.AllocsPerRun(100, func() {
		a.And(b)
	})

	if allocs != 0 {
		t.Error("Should NOT allocate", allocs)
	}

	if a.Get(1) != 0 {
		t.Error("Did not remove set 1 correctly", a, b)
	}

	if a.Get(27) != 1 {
		t.Error("Did not preserve 27 correctly", a, b)
	}

	if a.Get(25) == 1 {
		t.Error("Did not preserve 25 correctly", a, b)
	}

	if b.Get(1) == 1 {
		t.Error("Did not preserve b.1 correctly", a, b)
	}

}

func TestPopCount(t *testing.T) {
	b := New(100)
	b.Set(14)
	b.Set(25)
	b.Set(27)
	b.Set(30)
	b.Set(31)

	allocs := testing.AllocsPerRun(100, func() {
		b.PopCount()
	})

	if allocs != 0 {
		t.Error("Should NOT allocate", allocs)
	}

	if b.PopCount() != 5 {
		t.Error("PopCount was not correct: ", b.PopCount())
	}

	b2 := New(100)

	for i := 0; i < 100; i++ {
		b2.Set(uint(i))
		if b2.PopCount() != i+1 {
			t.Error("PopCount was not correct: ", b2.PopCount(), i)
		}
	}

}

func randomIndexes(b *testing.B, size uint) []uint {
	r := rand.New(rand.NewSource(0))
	indexes := make([]uint, 0, b.N)
	for i := 0; i < b.N; i++ {
		indexes = append(indexes, uint(r.Intn(int(size))))
	}
	return indexes
}

func BenchmarkSetBit(b *testing.B) {
	size := uint(300)
	bits := New(size)
	indexes := randomIndexes(b, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bits.Set(indexes[i])
	}
}

func BenchmarkGetBit(b *testing.B) {
	size := uint(300)
	bits := New(size)
	indexes := randomIndexes(b, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bits.Get(indexes[i])
	}
}

func BenchmarkAnd(b *testing.B) {
	size := uint(300)
	bits := New(size)
	bits2 := New(size)
	bits.Set(299)
	bits.Set(298)
	bits.Set(297)
	bits2.Set(297)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bits.And(bits2)
	}
}
