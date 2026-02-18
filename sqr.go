package mathutil

import (
	"math/big"
)

func (f *float) sqr() *float {
	if f == nil || f.n == nil {
		return f
	}

	f.n.Mul(f.n, f.n)
	f.fracBits *= 2

	if f.fracBits > f.maxFracBits {
		shift := f.fracBits - f.maxFracBits
		half := new(big.Int).Lsh(big.NewInt(1), uint(shift-1))
		f.n.Add(f.n, half)
		f.n.Rsh(f.n, uint(shift))
		f.fracBits = f.maxFracBits
	}

	return f
}
