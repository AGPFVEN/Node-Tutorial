package main

import "fmt"

func primitives_tutorial_() {
	//How booleans works
	var q bool = true
	fmt.Printf("%v, %T\n", q, q)

	//You can reassing them
	//Also reassign them using q logical operation (or not)
	q = 1 == 1
	fmt.Printf("%v, %T\n", q, q)

	//Obiously you can initialize q variable using q boolean operation (and use the standard way)
	w := 1 == 2
	fmt.Printf("%v, %T\n", w, w)

	//In go, q variable which has been initialized but not assigned will have q zero value (including booleans (it will be false))
	var e bool
	fmt.Printf("%v, %T\n", e, e)


	//How numbers work
	var r int8 = 1	//You can assign how many bites you want to store in an int
	fmt.Printf("%v, %T\n", r, r)

	var t uint		//You can initialize an unassinged int using uint (they can't be of 64 bits)
	fmt.Printf("%v, %T\n", t, t)

	//Basic arithmetic operations
	y := 10
	u := 3
	fmt.Println(y + u)
	fmt.Println(y - u)
	fmt.Println(y * u)
	fmt.Println(y / u)
	fmt.Println(y % u)

	//You cannot operate different types of intagers	
	//fmt.Println(r + u) This would raise the error "invalid operation (missmatched types)"

	//Floating point numbers
	i := 3.14 			//Standard way (This way will always use float64)
	i = 13.7e72			//We can use exponential notation (13.7 * 10⁷²)
	i = 2.1E14			//We can also use capital e to denote the same
	fmt.Printf("%v, %T\n", i, i)

	//Arithmetics with floating point numbers
	o := 10.2			//Floating point number
	p := 3.7			//Floating pint number
	fmt.Println(o + p)	//As we are operating two floating point numbers the result will be one floating point number 
	fmt.Println(o - p)	//(with all the problems that it has)
	fmt.Println(o * p)
	fmt.Println(o / p)

	//Text (Strings(utf-8), Runes(int32))
	a := "this is text"
	//a[2] = "u"										This line would raise an error because we cannot modify an string this way
	//Strings can be treated as arrays
	fmt.Printf("%v, %T\n", a[2], a[2]) 					//But this will return us the binary of the letter
	fmt.Printf("%v, %T\n", string(a[2]), string(a[2]))	//And this will return the character as a string

	//We can concatenate strings
	s := "This is also a text"
	fmt.Printf("%v, %T\n", a + s, a + s)

	//We can also use Runes, which is an alias of type int32
	d := 'd'
	fmt.Printf("%v, %T\n", d, d)

	var f rune = 'f' 									//This would be an strictly typed initialization of a rune
	fmt.Printf("%v, %T\n", f, f)
}