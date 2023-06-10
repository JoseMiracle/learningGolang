package main

import (
	"fmt"
	"strconv" // package for conversion

)

func main(){
	num := 3;
	fmt.Printf("The data type of the number %d before conversion is %T\n", num, num);
	
	int_to_str_num := strconv.Itoa(num); 
	fmt.Printf("The data type of the number %d after conversion is %T\n\n", num, int_to_str_num);



	str_num := "30";
	fmt.Printf("The data type of the number '%s' before conversion is %T\n", str_num, str_num);
	
	str_to_int_num,_ := strconv.Atoi(str_num); 
	fmt.Printf("The data type of the number '%s' after conversion is %T\n", str_num, str_to_int_num);

}