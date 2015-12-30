package enum

import (
	"strconv"
)

type Enum int

const NotFound = Enum(-1)

func (e Enum) String(t *Type) string {
	return t.String(e)
}

type Type struct {
	Name string
	n    []string
	m    map[string]Enum
}

func New(name string) *Type {
	return &Type{
		Name: name,
		n:    []string{},
		m:    make(map[string]Enum),
	}
}

func (t *Type) Iota(s string) Enum {
	e := Enum(len(t.n))
	t.n = append(t.n, s)
	t.m[s] = e
	return e
}

func (t *Type) String(e Enum) string {
	i := int(e)
	if i >= 0 && i < len(t.n) {
		return t.n[i]
	}
	return t.Name + "(" + strconv.Itoa(i) + ")"
}

func (t *Type) Lookup(s string) Enum {
	if e, ok := t.m[s]; ok {
		return e
	}
	return NotFound
}

func (t *Type) Names() []string {
	s := make([]string, len(t.n))
	copy(s, t.n)
	return s
}

func (t *Type) Enums() []Enum {
	values := make([]Enum, len(t.n))
	for i, s := range t.n {
		values[i] = t.m[s]
	}
	return values
}

func (t *Type) IsValid(s string) bool {
	_, ok := t.m[s]
	return ok
}

type TypeWithInterface struct {
	*Type
	i []interface{}
}

func NewWithInterface(name string) *TypeWithInterface {
	return &TypeWithInterface{
		Type: New(name),
	}
}

func (t *TypeWithInterface) Iota(s string, i interface{}) Enum {
	e := t.Type.Iota(s)
	t.i = append(t.i, i)
	return e
}

func (t *TypeWithInterface) LookupValue(s string) interface{} {
	if e, ok := t.m[s]; ok {
		return t.i[e]
	}
	return NotFound
}
