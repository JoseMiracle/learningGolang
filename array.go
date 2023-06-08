package main

import "fmt"

func main(){
	names := [5]string{"Miracle", "Joseph", "Joshua", "Bright", "Joycee"};

	for index,name:= range names{
		fmt.Printf("The name at index %d is %s\n", index, name);
	}
}