//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The frac package implements fractions.
package frac

import (
	"fmt"
    "errors"
)

//A Frac is a fraction.
type Frac struct {
	num, den uint64
	positive bool
}

var DivByZero = errors.New("Attempt to divide by zero.")

func abs(i int64) uint64 {
	if i >= 0 {
		return uint64(i)
	}
	return uint64(-i)
}

func gcd(x, y uint64) uint64 {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

//New() returns a *Frac unless den==0
func New(num, den int64) (*Frac, error) {
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

//f.String() satisfies fmt.Stringer
func (f *Frac) String() string {
	return fmt.Sprintf("%d/%d", f.Numerator(), f.Denominator())
}

func (f *Frac) simplify() {
	if f.num == 0 { //This is not strictly necessary, but it prevents
		f.den = 1 //f.String() from outputing "-0"
		return
	}
	common := gcd(f.num, f.den)
	f.num /= common
	f.den /= common
}

//f.Positive() returns true if f>0 and false if f<0. Behavior at f==0 is undefined.
func (f *Frac) Positive() bool {
	return f.positive
}

//f.Plus(other) returns f + other, leaving both unchanged.
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

//f.Negative() returns -f.
func (f *Frac) Negative() *Frac {
	return &Frac{f.num, f.den, !f.positive}
}

//f.Minus(other) returns f - other, leaving both unchanged.
func (f *Frac) Minus(other *Frac) *Frac {
	return f.Plus(other.Negative())
}

//f.Inverse() returns f^-1.
func (f *Frac) Inverse() *Frac {
	return &Frac{f.den, f.num, f.positive}
}

//f.Times(other) returns f * other, leaving both unchanged.
func (f *Frac) Times(other *Frac) *Frac {
	ret := &Frac{f.num * other.num, f.den * other.den, f.positive == other.positive}
	ret.simplify()
	return ret
}

//f.Divided(other) returns f ÷ other, leaving both unchanged.
func (f *Frac) Divided(other *Frac) (*Frac, error) {
	if other.num == 0 {
		return nil, DivByZero
	}
	return f.Times(other.Inverse()), nil
}

//f.Numerator() returns the (signed) numerator of f.
func (f *Frac) Numerator() int64 {
	ret := int64(f.num)
	if !f.positive {
		ret *= -1
	}
	return ret
}

//f.Denominator() returns the (always positive) denominator of f.
func (f *Frac) Denominator() int64 {
	return int64(f.den)
}

//f.Float64() returns a Float64 approximation of f.
func (f *Frac) Float64() float64 {
	ret := float64(f.num) / float64(f.den)
	if !f.positive {
		ret *= -1
	}
	return ret
}

//f.Mixed() returns a string like "2 1/3".
func (f *Frac) Mixed() string {
	if f.num < f.den {
		return f.String()
	}
	if f.den == 1 {
		return fmt.Sprint(f.Numerator())
	}
	if f.positive {
		return fmt.Sprintf("%d %d/%d", f.num/f.den, f.num%f.den, f.den)
	}
	return fmt.Sprintf("-%d %d/%d", f.num/f.den, f.num%f.den, f.den)
}
