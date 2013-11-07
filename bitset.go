/*
	A zero-allocation bitset library. After creation of a BitSet, all operations besides String() should not allocate.


	This package is similar in many ways to https://github.com/willf/bitset but is different in one important way: none of the functions perform an allocation except String().  This has profound impacts on the API and the implementation. Another feature which is vital for my use case is a fast way to iterate over the indexes of the bits set to 1.
*/
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

//create a new BitSet and return it. size is number of BITS.
//The bitset's backing array is allocated immediately.
func New(size uint) BitSet {
	byteLength := int(math.Ceil(float64(size) / 8.0))
	return &bitSet{size, make([]byte, byteLength)}
}

//sets a single bit to 1
func (b *bitSet) Set(index uint) {
	b.bits[index>>3] |= 1 << (index & 7)
}

//Gets a single bit's value. the return value will be 0 or 1
func (b *bitSet) Get(index uint) byte {
	result := b.bits[index>>3] >> (index & 7) & 1

	return result
}

//Sets a single bit to 0
func (b *bitSet) Clear(index uint) {
	b.bits[index>>3] &= ^(1 << (index & 7))
}

//logical ANDs with the receiver
func (b *bitSet) And(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] &= v
	}
}

// receiver &= ^ argument
func (b *bitSet) AndNot(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] &= ^v
	}
}

// receiver |= argument
func (b *bitSet) Or(o BitSet) {
	for i, v := range o.Bytes() {
		b.bits[i] |= v
	}
}

//Copies the bits from the argument to the receiver
func (b *bitSet) CopyFrom(o BitSet) {
	copy(b.bits, o.Bytes())
}

//Set all bits to 0
func (b *bitSet) Zero() {
	for i, _ := range b.bits {
		b.bits[i] = 0
	}
}

//Get the bytes that represent this BitSet. Does not allocate.
func (b *bitSet) Bytes() []byte {
	return b.bits
}

//ALLOCATES!  Creates a string representation of the BitSet
func (b *bitSet) String() string {
	s := ""
	for _, v := range b.bits {
		for i := uint(0); i <= 7; i++ {
			s += fmt.Sprint((v >> i) & 1)
		}
	}
	return s
}

//The size of the BitSet
func (b *bitSet) BitLen() uint {
	return b.size
}
