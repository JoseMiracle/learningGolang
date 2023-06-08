package main

import "fmt"

func main(){
	
	var number_of_names int;
	fmt.Printf("Enter the number of names you want to enter: ");
	fmt.Scanf("%d", &number_of_names);

	names := [5]string{} //Array to store name


	//This to get names from user and store in the slice names
	for i := 0; i < number_of_names; i++{
		num := i+1 ;
		fmt.Printf("Enter the name of person %d: ", num);
		
		var name string;
		fmt.Scanf("%s", &name);
		names[i] = name;
	} 
	
	fmt.Println("Below are the names of the users you entered \n");
	for index,name := range names{
		if(len(name) == 0){ 
			break
		}else{
			num := index+1;
				fmt.Printf("The name of person %d is %s\n", num, name);
		}
		
	}
}