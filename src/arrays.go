package main
import "fmt"

func main() {

	var numbers [3] string //Declaring a string array of size 3 and adding elements
	numbers[0] = "One"
	numbers[1] = "Two"
	numbers[2] = "Three"
	fmt.Println(numbers[1]) //prints Two
	fmt.Println(len(numbers)) //prints 3
	fmt.Println(numbers) // prints [One Two Three]

	directions := [...] int {1,2,3,4,5} // creating an integer array and the size of the array is defined by the number of elements
	fmt.Println(directions) //prints [1 2 3 4 5]
	fmt.Println(len(directions)) //prints 5

	//Executing the below commented statement prints invalid array index 5 (out of bounds for 5-element array)
	//fmt.Println(directions[5])


	// declaring array
	a := [5] string {"one", "two", "three", "four", "five"}
	fmt.Println("Array after creation:",a)

	var b [] string = a[1:4] //created a slice named b
	fmt.Println("Slice after creation:",b)

	b[0]="changed" // changed the slice data
	fmt.Println("Slice after modifying:",b)
	fmt.Println("Array after slice modification:",a)
}
