package bloom

import (
	"hash/fnv"
)

type BloomFilter struct {
	bitset []bool
	size   uint32
}

func New(size uint32) *BloomFilter {
	return &BloomFilter{
		bitset: make([]bool, size),
		size:   size,
	}
}

func (bf *BloomFilter) hash1(key string) uint32 {
	h := fnv.New32()
	h.Write([]byte(key))
	return h.Sum32() % bf.size
}

func (bf *BloomFilter) hash2(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32() % bf.size
}

func (bf *BloomFilter) Add(key string) {
	bf.bitset[bf.hash1(key)] = true
	bf.bitset[bf.hash2(key)] = true
}

func (bf *BloomFilter) MightContain(key string) bool {
	return bf.bitset[bf.hash1(key)] &&
		bf.bitset[bf.hash2(key)]
}
