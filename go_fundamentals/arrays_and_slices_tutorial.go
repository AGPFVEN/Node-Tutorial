package main

import (
	"fmt"
)

func main(){
	//Arrays
	grades := [3]int{100, 90, 80}					//This is the way to declare an array
	grades_improved := [...]int{100, 80, 70, 60}	//The three dots at the beginning to automate the number of items in the array

	fmt.Printf("Grades: %v, %T\n", grades, grades)
	fmt.Printf("Grades: %v\n", grades_improved)

	var students [3]string							//This is other way to initialize an array
	fmt.Printf("Students: %v\n", students)

	students[0] = "Lisa"							//And this would be the way to assign (or reassign) items in the array
	fmt.Printf("Students: %v\n", students)

	var identit_matrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}	//Multidimensional arrays
	fmt.Println(identit_matrix)

	q := grades 									//This will literally copy the array to new memory
	w := &grades 									//This will point to the memory of the array (pointer)
	fmt.Println(q, w)

	grades[0] = 10									//This will change grades and w, but not q
	fmt.Println(q, w)

	//Slices
	e := []int{1, 2, 3, 4, 5, 6}					//Nothing at the beginning to declare an slice
	r := e											//Slices only can point to each other
	r[1] = 5
	fmt.Println(e)
	fmt.Println(r)

	t := e[:]										//Slice all elements of the original slice
	y := e[1:4]										//Slicing at the finest (like Python), it also works with arrays
	fmt.Println(t, y)

	e = append(e, 7)
	fmt.Println(e)									//Slices don't have a fixed value for their entire life

	e = append(e, 8, 9, 10)							//append() is a variatic function (after the name of the array all will be appended)
	fmt.Println(e)
}