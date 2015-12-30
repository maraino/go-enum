package main

import (
	"fmt"

	"github.com/maraino/go-enum"
)

type Planet struct {
	Mass   float64
	Radius float64
}

var (
	days      = enum.New("Days")
	SUNDAY    = days.Iota("Sunday")
	MONDAY    = days.Iota("Monday")
	TUESDAY   = days.Iota("Tuesday")
	WEDNESDAY = days.Iota("Wednesday")
	THURSDAY  = days.Iota("Thursday")
	FRIDAY    = days.Iota("Friday")
	SATURDAY  = days.Iota("Saturday")

	planets = enum.NewWithInterface("Planets")
	MERCURY = planets.Iota("Mercury", Planet{3.303e+23, 2.4397e6})
	VENUS   = planets.Iota("Venus", Planet{4.869e+24, 6.0518e6})
	EARTH   = planets.Iota("Earth", Planet{5.976e+24, 6.37814e6})
	MARS    = planets.Iota("Mars", Planet{6.421e+23, 3.3972e6})
	JUPITER = planets.Iota("Jupiter", Planet{1.9e+27, 7.1492e7})
	SATURN  = planets.Iota("Saturn", Planet{5.688e+26, 6.0268e7})
	URANUS  = planets.Iota("Uranus", Planet{8.686e+25, 2.5559e7})
	NEPTUNE = planets.Iota("Neptune", Planet{1.024e+26, 2.4746e7})
)

func main() {
	fmt.Println("Days:")
	for _, v := range days.Enums() {
		fmt.Printf("%10s is %d\n", v.String(days), v)
	}

	fmt.Println()

	fmt.Println("Planets:")
	for _, s := range planets.Names() {
		planet := planets.LookupValue(s).(Planet)
		v := planets.Lookup(s)
		fmt.Printf("%9s is %d, mass: %e kg. and radius: %e m.\n", s, v, planet.Mass, planet.Radius)
	}
}
