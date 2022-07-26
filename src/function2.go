package main

import (
	"calculation"
	"fmt"
)
//import "calculation"  // import the package
//import "./calculation"  // import the package

//import (
//	// If you wan't to import a package with a different name, you can write the alias and then
//	// the package name like this:
//	calculation "./calculation"
//)
// not working yet


func main() {
	x,y := 15,10

	//calls the function calc with x and y an d gets sum, diff as output
	sum := calculation.Do_add(x,y)
	fmt.Println("Sum",sum)
	//fmt.Println("Diff",diff)


}