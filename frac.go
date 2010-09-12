//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The frac package implements high-precision fractions.
package frac

import (
	"os"
	"fmt"
)

type Frac struct {
	num, den uint
	positive bool
}

type error uint

const (
	DivByZero error = iota
)

func abs(i int) uint {
	if i>=0{
		return uint(i)
	}
	return uint(-i)
}

func (e error) String() string {
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
		num:      abs(num),
		den:      abs(den),
		positive: num == 0 || ((num > 0) == (den > 0)),
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

func (f *Frac) Positive() bool {
	return f.positive
}

func (f *Frac) Plus(other *Frac) *Frac {
	if f.den == other.den {
		ret := new(Frac)
		ret.den = f.den
		switch {
		case f.positive && other.positive:
			ret.positive = true
			ret.num = f.num + other.num
		case !f.positive && !other.positive:
			ret.positive = false
			ret.num = f.num + other.num
		case f.num == other.num:
			ret.num = 0
		case f.num > other.num:
			ret.positive = f.positive
			ret.num = f.num - other.num
		case f.num < other.num:
			ret.positive = other.positive
			ret.num = other.num - f.num
		}
		ret.simplify()
		return ret
	}
	return (&Frac{f.num * other.den, f.den * other.den, f.positive}).Plus(&Frac{other.num * f.den, other.den * f.den, other.positive})
}

func (f *Frac) Negative() *Frac {
	return &Frac{f.num, f.den, !f.positive}
}

func (f *Frac) Minus(other *Frac) *Frac {
	return f.Plus(other.Negative())
}

func (f *Frac) Inverse() *Frac {
	return &Frac{f.den, f.num, f.positive}
}

func (f *Frac) Times(other *Frac) *Frac {
	ret := &Frac{f.num * other.num, f.den * other.den, f.positive == other.positive}
	ret.simplify()
	return ret
}

func (f *Frac) Divided(other *Frac) (*Frac, os.Error) {
	if other.num == 0 {
		return nil, DivByZero
	}
	return f.Times(other.Inverse()), nil
}

func (f *Frac) Numerator() int {
	ret := int(f.num)
	if !f.positive {
		ret *= -1
	}
	return ret
}

func (f *Frac) Denominator() int {
	return int(f.den)
}

func (f *Frac) Float() float {
	ret := float(f.num) / float(f.den)
	if !f.positive {
		ret *= -1
	}
	return ret
}

func (f *Frac) Mixed() string {
	if f.positive {
		return fmt.Sprintf("%d %d/%d", f.num/f.den, f.num%f.den, f.den)
	}
	return fmt.Sprintf("-%d %d/%d", f.num/f.den, f.num%f.den, f.den)
}
