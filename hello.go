package main


import (

	"fmt"
	"os"
)


func main() {

	fmt.Println("Lubuntu Hello World")
	fmt.Println(os.Getenv("USER"),"-- Lubuntu current user")

}
