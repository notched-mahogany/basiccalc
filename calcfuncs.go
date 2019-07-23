//This is a basic calculator. It can only run addition, subtraction, division and multiplication
//Current bugs:
//	"For loops being go
//	"Splice issues with the count I'm calling being outside the index range  
package main


import (
	"fmt"
)

func fix() {
	fmt.Print("\nWhat number are you adding?\n>>")
	fmt.Scanln(&b)
	lst = append(lst, b)
	swing+=1
}

func selection(){
	var y = int(lst[count]-1)
		if choice == "A"{
			x := x+y
			_=x
		}
		if choice == "S"{
			x := x-y
			_=x
		}
		if choice == "D"{
			x := x/y
			_=x
		}
		if choice == "M"{
			x := x*y
			_=x
		}
		count+=1
}

func main() {
	fmt.Print("*************Calculator*************\n\nWhat operator would you like to use?\n\nAdd (input 'A')\nSubtract (input 'S')\nMultiply (input 'M')\nDivide (input 'D')\n\nTo Quit press 'Q'\n>>")
	var choice string
	fmt.Scanln(&choice)
	if choice != "Q" {
		var number int
		fmt.Print("How many numbers are you going to input?\n>>")
		fmt.Scanln(&number)
		var lst []int
		var b int = 1
		var swing int = 1
		fmt.Print("What is the starting number?\n>>")
		fmt.Scan(&b)
		lst = append(lst, b)
		for ; swing <= number-1 ;{
			fix()
		}
		var x = int(lst[0])
		var count int = 1
		for ; count <= swing; {
			selection()
		}
		fmt.Print("The inputed values work out to ",x)
	}
}
