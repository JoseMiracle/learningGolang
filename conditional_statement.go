package main
import "fmt"

func main(){
	num1 := 3;
	num2 := 5;
	
	if(num1 > num2){
		fmt.Printf("%v is greater than %v \n", num1, num2);
	}else if(num2 > num1){
		fmt.Printf("%v is greater than %v\n", num2, num1);
	}
}