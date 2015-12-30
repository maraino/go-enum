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
	v    []string
	m    map[string]Enum
}

func New(name string) *Type {
	return &Type{
		Name: name,
		v:    []string{},
		m:    make(map[string]Enum),
	}
}

func (t *Type) Iota(s string) Enum {
	e := Enum(len(t.v))
	t.v = append(t.v, s)
	t.m[s] = e
	return e
}

func (t *Type) String(e Enum) string {
	i := int(e)
	if i >= 0 && i < len(t.v) {
		return t.v[i]
	}
	return t.Name + "(" + strconv.Itoa(i) + ")"
}

func (t *Type) Lookup(s string) Enum {
	if e, ok := t.m[s]; ok {
		return e
	}
	return NotFound
}

func (t *Type) Values() []Enum {
	values := make([]Enum, len(t.v))
	i := 0
	for _, v := range t.m {
		values[i] = v
		i++
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
	e := Enum(len(t.v))
	t.v = append(t.v, s)
	t.i = append(t.i, i)
	t.m[s] = e
	return e
}

func (t *TypeWithInterface) LookupValue(s string) interface{} {
	if e, ok := t.m[s]; ok {
		return t.i[e]
	}
	return NotFound
}
