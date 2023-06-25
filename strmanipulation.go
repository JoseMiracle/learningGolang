package main

import (
	"fmt"
	"strings" // package for string maninpulation
)	


func main(){
	name := "Miracle";
	
	is_name_contain_yi := strings.Contains(name, "le");
	fmt.Printf("%t\n", is_name_contain_yi);	

	is_prefix_Mir := strings.HasPrefix(name, "Mir");
	fmt.Printf("%t\n", is_prefix_Mir);

	words := []string{"Golang", "is", "a", "good" , "programming" , "language."};
	join_words := strings.Join(words, " ");
	fmt.Printf("%s\n", join_words);

}