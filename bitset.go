package bitset

import "math"
import "fmt"

type BitSet interface {
	//set the bit at index i to 1
	Set(i uint)

	//set the bit at index i to 0
	Clear(i uint)

	//gets the bit at index i
	Get(i uint) byte

	// a &= b
	And(b BitSet)

	// a ^= b
	AndNot(b BitSet)

	// a |= b
	Or(b BitSet)

	// a = 0
	Zero()

	// a = b
	CopyFrom(b BitSet)

	//zero-copy []byte access
	Bytes() []byte

	//zero-padded string representation. ALLOCATES!
	String() string

	BitLen() uint
}

type bitSet struct {
	size uint
	bits []byte
}

//create a new BitSet and return it.
//The bitset's backing array is allocated immediately.
func New(size uint) BitSet {
	byteLength := int(math.Ceil(float64(size) / 8.0))
	return &bitSet{size, make([]byte, byteLength)}
}

func (b *bitSet) Set(index uint) {
	b.bits[index>>3] |= 1 << (index & 7)
}

func (b *bitSet) Get(index uint) byte {
	result := b.bits[index>>3] >> (index & 7) & 1

	return result
}

func (b *bitSet) Clear(index uint) {
	b.bits[index>>3] &= ^(1 << (index & 7))
}

func (b *bitSet) And(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] &= v
	}
}

func (b *bitSet) AndNot(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] &= ^v
	}
}

func (b *bitSet) Or(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] |= v
	}
}

func (b *bitSet) CopyFrom(o BitSet) {
	copy(b.bits, o.Bytes())
}

func (b *bitSet) Zero() {
	for i, _ := range b.bits {
		b.bits[i] = 0
	}
}

func (b *bitSet) Bytes() []byte {
	return b.bits
}

func (b *bitSet) String() string {
	s := ""
	for _, v := range b.bits {
		for i := uint(0); i <= 7; i++ {
			s += fmt.Sprint((v >> i) & 1)
		}
	}
	return s
}

func (b *bitSet) BitLen() uint {
	return b.size
}
