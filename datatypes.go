package main
import "fmt"

// TODO: more data types will be added  
func main(){
	number := 32;
	fmt.Printf("The data type of %v is %T\n", number, number);
	
	number_1 := -3.5;
	fmt.Printf("The data type of %v is %T\n", number_1, number_1);

	number_2 := 555555555577777777; //I have to store in int64 bcos the value can't be stored datatype int32
	fmt.Printf("The data type of %v is %T\n", number_2, number_2);

	decimal_number := -333333335.5;
	fmt.Printf("The datatype of %v is %T\n", decimal_number,decimal_number);

	name := "Miracle";
	fmt.Printf("The data type of %v is %T\n", name, name);

	isTrue := true;
	fmt.Printf("The data type of %v is %T\n", isTrue, isTrue);
	

}