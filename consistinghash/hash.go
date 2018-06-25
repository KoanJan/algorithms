package consistinghash

import "crypto/sha256"

func Hash(src string) uint32 {
	ba := sha256.Sum256([]byte(src))
	sum := uint32(0)
	for i := 0; i < 32; i++ {
		sum |= uint32(ba[i])
		sum <<= 1
	}
	return sum
}
