package main

import "fmt"

func main(){
	// This is an array-like data structure that we don't specify the number of element it has to store.
	
	array_of_numbers := []int{};

	// Below we ask the user the number of numbers they want to enter and store.
	var number int;

	fmt.Printf("How many number do you want to enter ?: ");
	fmt.Scanf("%d", &number);

	index := 0
	for index < number{
		var num int;
		fmt.Printf("Enter the number for index %d: ", index);
		fmt.Scanf("%d", &num);
		array_of_numbers = append(array_of_numbers, num)
		index++;
	}

	for index, each_number := range array_of_numbers{
		fmt.Printf("The value at *index %d* is %d \n", index, each_number);
	}
	
}