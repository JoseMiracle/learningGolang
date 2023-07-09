package main
import (
	"fmt"
	"strconv"
)

func main(){
	_, err := strconv.Atoi("23b")
	if err != nil{
		fmt.Println("error occured while running the code...")
	}
}
