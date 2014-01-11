// Package worgen implements generating pronounceable words from integers.
package worgen

import (
	"hash/crc64"
	"math/big"
	"math/rand"
)

var (
	consonants = []rune{
		'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm',
		'n', 'p', 'q', 'r', 's', 't', 'v', 'y', 'z',
	}
	vovels = []rune{'a', 'o', 'u', 'e', 'i'}
	finishers = []rune{
		'b', 'c', 'd', 'f', 'g', 'h', 'k', 'l', 'm',
		'n', 'p', 'r', 's', 't', 'v', 'w', 'y', 'z',
		'\'',
	}
)

func genran(ran *rand.Rand) (s string) {
	s += string(consonants[ran.Intn(len(consonants))])
	s += string(vovels[ran.Intn(len(vovels))])
	s += string(finishers[ran.Intn(len(finishers))])
	return s
}

// Wor generates string of count words for 64 bit integer n.
func Wor(n int64, count int) (s string) {
	ran := rand.New(rand.NewSource(int64(n)))
	for i := 0; i < count; i++ {
		s += genran(ran)
		s += genran(ran)
		s += "-"
	}
	return s[:len(s)-1]
}

// Wor16 generates string of one word for 16 bit integer n.
func Wor16(n int16) string {
	return Wor(int64(n), 1)
}

// Wor32 generates string of two words for 32 bit integer n.
func Wor32(n int32) string {
	return Wor(int64(n), 2)
}

// Wor64 generates string of four words for 32 bit integer n.
func Wor64(n int64) string {
	return Wor(int64(n), 4)
}

// WorBig generates string for multi-precision integer n, count of words
// calculated as:
//     (bitLength(n) / 16) rounded towards infinity
func WorBig(n *big.Int) string {
	if n.BitLen() < 64 {
		return Wor64(n.Int64())
	} else {
		sum := crc64.Checksum(n.Bytes(), crc64.MakeTable(crc64.ECMA))
		return Wor64(int64(sum))
	}
}
