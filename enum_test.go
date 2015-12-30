package enum

import (
	"reflect"

	"testing"
)

var (
	e = New("MyEnum")
	A = e.Iota("A") // 0
	B = e.Iota("B") // 1

	ei = NewWithInterface("MyEnumWithInterface")
	AA = ei.Iota("AA", []string{"A", "A"})
	BB = ei.Iota("BB", []string{"B", "B"})
)

func TestString(t *testing.T) {
	// String enums
	if A.String(e) != "A" {
		t.Error()
	}
	if B.String(e) != "B" {
		t.Error()
	}
	if e.String(A) != "A" {
		t.Error()
	}
	if e.String(B) != "B" {
		t.Error()
	}
	if Enum(10).String(e) != "MyEnum(10)" {
		t.Error()
	}
	// With inteerface
	if AA.String(ei.Type) != "AA" {
		t.Error()
	}
	if BB.String(ei.Type) != "BB" {
		t.Error()
	}
	if ei.String(AA) != "AA" {
		t.Error()
	}
	if ei.String(BB) != "BB" {
		t.Error()
	}
	if Enum(10).String(ei.Type) != "MyEnumWithInterface(10)" {
		t.Error()
	}
}

func TestLookup(t *testing.T) {
	// String enums
	if e.Lookup("A") != A {
		t.Error()
	}
	if e.Lookup("B") != B {
		t.Error()
	}
	if e.Lookup("C") != NotFound {
		t.Error()
	}
	// With inteerface
	if ei.Lookup("AA") != AA {
		t.Error()
	}
	if ei.Lookup("BB") != BB {
		t.Error()
	}
	if ei.Lookup("CC") != NotFound {
		t.Error()
	}
}

func TestLookupValue(t *testing.T) {
	if !reflect.DeepEqual(ei.LookupValue("AA"), []string{"A", "A"}) {
		t.Error()
	}
	if !reflect.DeepEqual(ei.LookupValue("BB"), []string{"B", "B"}) {
		t.Error()
	}
	if ei.LookupValue("CC") != NotFound {
		t.Error()
	}
}

func TestValues(t *testing.T) {
	// String enums
	if !reflect.DeepEqual(e.Values(), []Enum{A, B}) {
		t.Error()
	}
	// With interface
	if !reflect.DeepEqual(ei.Values(), []Enum{AA, BB}) {
		t.Error()
	}
}

func TestValid(t *testing.T) {
	// String enums
	if !e.IsValid("A") {
		t.Error()
	}
	if !e.IsValid("B") {
		t.Error()
	}
	if e.IsValid("C") {
		t.Error()
	}
	// With interface
	if !ei.IsValid("AA") {
		t.Error()
	}
	if !ei.IsValid("BB") {
		t.Error()
	}
	if ei.IsValid("CC") {
		t.Error()
	}
}
