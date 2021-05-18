package CryptoUtils

import "hash/fnv"

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

//
// func I64tob(val uint64) []byte {
// 	r := make([]byte, 8)
// 	for i := uint64(0); i < 8; i++ {
// 		r[i] = byte((val >> (i * 8)) & 0xff)
// 	}
// 	return r
// }
//
// func Btoi64(val []byte) uint64 {
// 	r := uint64(0)
// 	for i := uint64(0); i < 8; i++ {
// 		r |= uint64(val[i]) << (8 * i)
// 	}
// 	return r
// }
//
// func I32tob(val uint32) []byte {
// 	r := make([]byte, 4)
// 	for i := uint32(0); i < 4; i++ {
// 		r[i] = byte((val >> (8 * i)) & 0xff)
// 	}
// 	return r
// }
//
// func Btoi32(val []byte) uint32 {
// 	r := uint32(0)
// 	for i := uint32(0); i < 4; i++ {
// 		r |= uint32(val[i]) << (8 * i)
// 	}
// 	return r
// }
