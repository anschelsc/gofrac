package frac

import (
	"testing"
)

func TestSanity(t *testing.T) {
	if 0 != 0 {
		t.Errorf("0 is not equal to 0, so the universe is meaningless")
	}
}

func TestNew(t *testing.T) {
	f, err := New(3, 4)
	if err != nil {
		t.Errorf("New(3, 4) returned error %v, expected nil", err)
	}
	if num := f.Numerator(); num != 3 {
		t.Errorf("Expected numerator 3, got %v", num)
	}
	if den := f.Denominator(); den != 4 {
		t.Errorf("Expected denominator 4, got %v", den)
	}
	if !f.Positive() {
		t.Errorf("Expected positive, got negative")
	}

	f, err = New(7, 14)
	if err != nil {
		t.Errorf("New(7, 14) returned error %v, expected nil", err)
	}
	if num := f.Numerator(); num != 1 {
		t.Errorf("Expected numerator 1, got %v", num)
	}
	if den := f.Denominator(); den != 2 {
		t.Errorf("Expected denominator 2, got %v", den)
	}
	if !f.Positive() {
		t.Errorf("Expected positive, got negative")
	}

	f, err = New(14, 7)
	if err != nil {
		t.Errorf("New(14, 7) returned error %v, expected nil", err)
	}
	if num := f.Numerator(); num != 2 {
		t.Errorf("Expected numerator 2, got %v", num)
	}
	if den := f.Denominator(); den != 1 {
		t.Errorf("Expected denominator 1, got %v", den)
	}
	if !f.Positive() {
		t.Errorf("Expected positive, got negative")
	}

	f, err = New(13, 0)
	if err.(error) != DivByZero {
		t.Errorf("Expected DivByZero from New(13, 0), got %v", err)
	}

	f, err = New(-20, 15)
	if err != nil {
		t.Errorf("New(-20, 15) returned error %v, expected nil", err)
	}
	if num := f.Numerator(); num != -4 {
		t.Errorf("Expected numerator -4, got %v", num)
	}
	if den := f.Denominator(); den != 3 {
		t.Errorf("Expected denominator 3, got %v", den)
	}
	if f.Positive() {
		t.Errorf("Expected negative, got positive")
	}

	f, err = New(3, -3)
	if err != nil {
		t.Errorf("New(3, -3) returned error %v, expected nil", err)
	}
	if num := f.Numerator(); num != -1 {
		t.Errorf("Expected numerator -1, got %v", num)
	}
	if den := f.Denominator(); den != 1 {
		t.Errorf("Expected denominator 1, got %v", den)
	}
	if f.Positive() {
		t.Errorf("Expected negative, got positive")
	}
}

func TestMult(t *testing.T) {
	positive, _ := New(1, 2)
	negative, _ := New(-1, 2)
	if s := positive.Times(positive).String(); s != "1/4" {
		t.Errorf("(1/2) * (1/2) = %v", s)
	}
	if s := negative.Times(positive).String(); s != "-1/4" {
		t.Errorf("(-1/2) * (1/2) = %v", s)
	}
	if s := positive.Times(negative).String(); s != "-1/4" {
		t.Errorf("(1/2) * (-1/2) = %v", s)
	}
	if s := negative.Times(negative).String(); s != "1/4" {
		t.Errorf("(-1/2) * (-1/2) = %v", s)
	}

	third, _ := New(1, 3)
	fourFifths, _ := New(4, 5)
	nine, _ := New(9, 1)
	if s := third.Times(fourFifths).String(); s != "4/15" {
		t.Errorf("(1/3) * (4/5) = %v", s)
	}
	if s := fourFifths.Times(nine).String(); s != "36/5" {
		t.Errorf("(4/5) * (9/1) = %v", s)
	}
	if s := third.Times(nine).String(); s != "3/1" {
		t.Errorf("(1/3) * (9/1) = %v", s)
	}
}

func TestAddAndNeg(t *testing.T) {
	half, _ := New(1, 2)
	third, _ := New(1, 3)
	minusHalf := half.Negative()
	minusThird := third.Negative()
	if s := minusHalf.String(); s != "-1/2" {
		t.Errorf("-(1/2) = %v", s)
	}
	if s := minusThird.String(); s != "-1/3" {
		t.Errorf("-(1/3) = %v", s)
	}
	if s := half.Plus(third).String(); s != "5/6" {
		t.Errorf("(1/2) + (1/3) = %v", s)
	}
	if s := half.Plus(minusThird).String(); s != "1/6" {
		t.Errorf("(1/2) - (1/3) = %v", s)
	}
	if s := minusHalf.Plus(third).String(); s != "-1/6" {
		t.Errorf("-(1/2) + (1/3) = %v", s)
	}
	if s := minusHalf.Plus(minusThird).String(); s != "-5/6" {
		t.Errorf("-(1/2) - (1/3) = %v", s)
	}
}

func TestNumAndDen(t *testing.T) {
	half, _ := New(1, 2)
	minusHalf, _ := New(1, -2)
	if n := half.Numerator(); n != 1 {
		t.Errorf("half.Numerator() == %v", n)
	}
	if n := half.Denominator(); n != 2 {
		t.Errorf("half.Denominator() == %v", n)
	}
	if n := minusHalf.Numerator(); n != -1 {
		t.Errorf("minusHalf.Numerator() == %v", n)
	}
	if n := minusHalf.Denominator(); n != 2 {
		t.Errorf("minusHalf.Denominator() == %v", n)
	}
}

func TestInverse(t *testing.T) {
	seventeen, _ := New(17, 1)
	threeSeventeenths, _ := New(3, 17)
	if s := seventeen.Inverse().String(); s != "1/17" {
		t.Errorf("(1/17)^-1 = %v", s)
	}
	if s := threeSeventeenths.Inverse().String(); s != "17/3" {
		t.Errorf("(3/17)^-1 = %v", s)
	}
}

func TestMixed(t *testing.T) {
	three, _ := New(3, 1)
	third, _ := New(1, 3)
	fourThirds, _ := New(4, 3)
	if m := three.Mixed(); m != "3" {
		t.Errorf("%v = %v", three, m)
	}
	if m := third.Mixed(); m != "1/3" {
		t.Errorf("%v = %v", third, m)
	}
	if m := fourThirds.Mixed(); m != "1 1/3" {
		t.Errorf("%v = %v", fourThirds, m)
	}
	if m := three.Negative().Mixed(); m != "-3" {
		t.Errorf("%v = %v", three.Negative(), m)
	}
	if m := third.Negative().Mixed(); m != "-1/3" {
		t.Errorf("%v = %v", third.Negative(), m)
	}
	if m := fourThirds.Negative().Mixed(); m != "-1 1/3" {
		t.Errorf("%v = %v", fourThirds.Negative(), m)
	}
}
