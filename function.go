package main

import (
	"fmt"
)

// func sum(number1, number2 int) int {
// 	return number1 + number2;
// }

func concatenate(str1, str2, string) string {
	return str1 + str2;
}

func main(){
	// number1 := 6
	// number2 := 7
	// result := sum(number1, number2);
	// fmt.Printf("%d", result);

	
	var str1 string = "Go";
	var str2 string = "lang";
	joined_string := concatenate(str1, str2);
	fmt.Printf("The conctenate string is: %s", joined_string);
}
