bitset
======

This package is similar in many ways to https://github.com/willf/bitset but is different in one important way: none of the functions perform an allocation except String().  This has profound impacts on the API and the implementation. Another feature which is vital for my use case is a fast way to iterate over the indexes of the bits set to 1.


