package main
import (
	"fmt"
	"strconv"
)

func main(){
	_, err := strconv.Atoi("23b2")
	if err != nil{
		fmt.Println("error occured while running the code...")
	}
}
