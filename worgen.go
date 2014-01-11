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

func Wor(id int64, count int) (s string) {
	ran := rand.New(rand.NewSource(int64(id)))
	for i := 0; i < count; i++ {
		s += genran(ran)
		s += genran(ran)
		s += "-"
	}
	return s[:len(s)-1]
}

func Wor16(n int16) string {
	return Wor(int64(n), 1)
}

func Wor32(n int32) string {
	return Wor(int64(n), 2)
}

func Wor64(n int64) string {
	return Wor(int64(n), 4)
}

func WorBig(n *big.Int) string {
	if n.BitLen() < 64 {
		return Wor64(n.Int64())
	} else {
		sum := crc64.Checksum(n.Bytes(), crc64.MakeTable(crc64.ECMA))
		return Wor64(int64(sum))
	}
}
