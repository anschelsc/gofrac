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

	f,err=New(14,7)
	if err!=nil{
		t.Errorf("New(14, 7) returned error %v, expected nil", err)
	}
	if num:=f.Numerator(); num!=2{
		t.Errorf("Expected numerator 2, got %v", num)
	}
	if den:=f.Denominator(); den!=1{
		t.Errorf("Expected denominator 1, got %v", den)
	}
	if !f.Positive() {
		t.Errorf("Expected positive, got negative")
	}

	f, err = New(13, 0)
	if err.(Error) != DivByZero {
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
}
