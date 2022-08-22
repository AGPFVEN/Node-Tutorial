package main

import (
	"fmt"
	//"math"
)

//iota constants
const (
	u = iota
	i										//This syntax is the same as:
	o2										//o2 = iota
)

func constants_tutorial_() {
	//iota 
	fmt.Printf("%v, %T\n", u, u)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", o2, o2)

	//Obviously a constant cannot be reassigned
	//A constant can be shadowed
	const u int = 5
	fmt.Printf("%v, %T\n", u, u)

	const q int = 1							//Typed constant, as a typed variable
	fmt.Printf("%v, %T\n", q, q)

	const r = 16							//Constants can non-strictly typed
	var t int16 = 27
	var y int64 = 27
	fmt.Printf("%v, %T\n", r + t, r + t)	//Constants can be implicitly converted
	fmt.Printf("%v, %T\n", r + y, r + y)
	//fmt.Printf("%v, %T\n", t + y, t + y)	//Variables cannot

	//const w = math.sin(0.9)	A constant cannot be assigned in runtime, it must be declared and assigned at the moment it is created
}