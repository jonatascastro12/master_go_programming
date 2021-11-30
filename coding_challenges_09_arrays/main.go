package main

import "fmt"

func main() {
	ex3()
}

func ex1() {
	var cities [2]string
	fmt.Println("cities", cities)

	grades := [3]float64{4, 5}
	fmt.Println("grades", grades)

	taskDone := [...]bool{}
	fmt.Println("taskDone", taskDone)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for _, val := range grades {
		fmt.Println(val)
	}
}

func ex2() {
	nums := [...]int{30, -1, -6, 90, -6}
	var count int

	for _, num := range nums {
		if num > 0 && num%2 == 0 {
			count++
		}
	}

	fmt.Println("Number of positive even numbers from \"nums\" array: ", count)
}

func ex3(){
	myArray := [4]float64{1.2, 5.6}
 
    myArray[2] = 6
  
    a := 10

    myArray[0] = float64(a)
 
    myArray[3] = 10.10
 
    fmt.Println(myArray)
}

