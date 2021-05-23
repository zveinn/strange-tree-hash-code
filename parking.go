package main

import (
	"hash/crc32"
	"strconv"
)

func GetKey(s string) (key int16) {
	for i, v := range []byte(s) {
		key += int16(int(v) * i)
	}
	return
}

func strToBinary(s string, base int) []byte {

	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), base)
	}

	return b
}
func Hash(s string) uint32 {
	v := crc32.ChecksumIEEE([]byte(s))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
