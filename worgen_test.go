package worgen

import (
	"math"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func stringCmp(a, b string, t *testing.T) {
	if a != b {
		t.Errorf("strings don't match: %q != %q", a, b)
	}
}

func TestDeterministic(t *testing.T) {
	a := Wor32(9)
	b := Wor32(9)
	t.Logf("a == %q", a)
	t.Logf("b == %q", b)
	stringCmp(a, b, t)
}

func TestBig(t *testing.T) {
	a := Wor64(9)
	b := WorBig(big.NewInt(9))
	t.Logf("a == %q", a)
	t.Logf("b == %q", b)
	stringCmp(a, b, t)
}

func TestWorgen(t *testing.T) {
	t.Logf("%20d\t%s", 0, Wor64(0))
	for i := 0; i < 4; i++ {
		n := rand.Int63()
		t.Logf("%20d\t%s", n, Wor64(n))
	}
}

func BenchmarkWorgen16(b *testing.B) {
	var n int16
	for i := 0; i < b.N; i++ {
		Wor16(n)
		n++
	}
}

func BenchmarkWorgen32(b *testing.B) {
	var n int32
	for i := 0; i < b.N; i++ {
		Wor32(n)
		n++
	}
}

func BenchmarkWorgen64(b *testing.B) {
	var n int64
	for i := 0; i < b.N; i++ {
		Wor64(n)
		n++
	}
}

func BenchmarkWorgenBig(b *testing.B) {
	var (
		n = big.NewInt(math.MaxInt64)
		one = big.NewInt(1)
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WorBig(n)
		n.Add(n, one)
	}
}
