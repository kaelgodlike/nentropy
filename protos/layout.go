package protos

import (
	"hash/fnv"
	"math/bits"
)

func Cale_bits_of(num int) int {
	return bits.Len(uint(num))
}

func Calc_pg_masks(pg_num int) int {
	pg_num_mask := 1<<uint(Cale_bits_of(pg_num-1)) - 1
	return pg_num_mask
}

func Calc_origin_pg(pg_id, old_pg_num int) int {
	old_pg_mask := Calc_pg_masks(old_pg_num)
	return pg_id & old_pg_mask
}

func Nentropy_str_hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Nentropy_stable_mod(x, b, bmask int) int {
	if (b & bmask) < b {
		return x & bmask
	} else {
		return x & (bmask >> 1)
	}
}
