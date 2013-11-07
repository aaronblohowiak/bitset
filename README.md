bitset
======

This package is similar in many ways to https://github.com/willf/bitset but is different in one important way: none of the functions perform an allocation except String().  This has profound impacts on the API and the implementation. Another feature which is vital for my use case is a fast way to iterate over the indexes of the bits set to 1.


This package is faster than using big.Int because setting a single bit in big.Int will cause all of the bytes that back it to be copied :( Additionally, because BitSet does not grow, it does not have to perform double-bounds-checking (check if needs to grow, and then the bounds check for the actual slice access.)

Docs: http://godoc.org/github.com/aaronblohowiak/bitset

