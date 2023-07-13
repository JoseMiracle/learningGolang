package main

import (
	"fmt"
)

func outputNumbers(nums...int){
	
	for i:=0; i < len(nums); i++{
	
		fmt.Println(nums[i])	
	}
	
}

func main(){
	outputNumbers(1,2,3,4);
	
}