package glib

import (
	"math/big"
)

// Uint64CheckedAdd adds two uint64s together and returns 0, true if the result won't fit in a uint64
func Uint64CheckedAdd(a uint64, b uint64) (val uint64, overflow bool) {
	aBig := big.NewInt(0)
	bBig := big.NewInt(0)
	aBig.SetUint64(a)
	bBig.SetUint64(b)
	aBig.Add(aBig, bBig)
	if !aBig.IsUint64() {
		return 0, true
	}
	return aBig.Uint64(), false
}

// Uint64CheckedMul adds two uint64s together and returns 0, true if the result won't fit in a uint64
func Uint64CheckedMul(a uint64, b uint64) (val uint64, overflow bool) {
	aBig := big.NewInt(0)
	bBig := big.NewInt(0)
	aBig.SetUint64(a)
	bBig.SetUint64(b)
	aBig.Mul(aBig, bBig)
	if !aBig.IsUint64() {
		return 0, true
	}
	return aBig.Uint64(), false
}
