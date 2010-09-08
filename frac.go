package frac

import (
	"os"
	"fmt"
)

type Frac struct {
	num, den uint
	positive bool
}

type Error uint

const (
	DivByZero Error = iota
)

func (e Error) String() string {
	switch e {
	case DivByZero:
		return "Attempt to divide by zero."
	}
	return "Unknown error."
}

func gcd(x, y uint) uint {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func New(num, den int) (*Frac, os.Error) {
	if den == 0 {
		return nil, DivByZero
	}
	f := &Frac{
		num:      uint(num),
		den:      uint(den),
		positive: num == 0 || (num > 0 && den > 0) || (num < 0 && den < 0),
	}
	f.simplify()
	return f, nil
}

func (f *Frac) String() string {
	if f.positive {
		return fmt.Sprintf("%d/%d", f.num, f.den)
	}
	return fmt.Sprintf("-%d/%d", f.num, f.den)
}

func (f *Frac) simplify() {
	if f.num == 0 {
		f.den = 1
		return
	}
	common := gcd(f.num, f.den)
	f.num /= common
	f.den /= common
}
