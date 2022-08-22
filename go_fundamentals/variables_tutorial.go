package main

import (
	"fmt"
	"strconv"
)

//If a variable is declared but not used it will raise a compile error

//There are only three scopes in go
//	Block scope: 	only visible for the function
//	Package scope:	only visible for the package (file)
//	Global scope:	globally visble (you need to declare the variable at package level with uppercase)

//It is possible to declare variables outside functions at "package-level"
var t int = 3 //But it needs to be declared like the 1st or 2nd way

//It is also possible to declare more than one variable at once
var(
	actorName string = "Elisabeth Sladen"
	companion string = "Sarah Jane Smith"
	doctorNumber int = 3
	season int = 11
)

//Shadowing: if you declare a variable at package-level
//and you redeclare it inside a function it will take the value of the innermost scope (the function)
var o = 0

func variables_tutorial_go() {
	//Three ways to declare variables

	var i int     //This way is used when you "don't know an initial value for the variable"
	i = 1

	var k int = 3 //This way is used when you need more control

	j := 2        //This is the "standard" way

	//Println is to use a simple print line
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	//Printf allows to provide format strings
	//"%v" is the value of the variable
	//"%T" is the type of the variable
	//"\n" is newline
	fmt.Printf("%v, %T, \n", i, i)

	var w float32 = 4
	fmt.Printf("%v, %T, \n", w, w)

	v := 5. //This is an abrebiation for the initialization of a variable of type float64
	fmt.Printf("%v, %T, \n", v, v)

	o := 7
	fmt.Printf("%v, %T, \n", o, o)

	//It is also possible to convert varibles, but you need a new vaiable for that
	var a float32
	a = float32(o)
	fmt.Printf("%v, %T, \n", a, a)

	//But if you want to convert something to strings you need to import the package "strconv"
	//Then use the function strconv.Itoa()
	var b string
	b = strconv.Itoa(o)
	fmt.Printf("%v, %T, \n", b, b)
}