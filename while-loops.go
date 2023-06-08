package main

import "fmt"


func main(){
	num:=1;
	for{
		if num > 5{
			break;
		}else{
			fmt.Println(num);
			num++;
		}
	}
}