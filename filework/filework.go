package filework

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func Write_to_file(filename string, amount int) {

	str_amount := strconv.Itoa(amount)
	err := os.WriteFile(filename, []byte(str_amount), 0644)

	if err != nil {
		fmt.Println("Unable to write file")
		log.Fatal(err)
	} else {
		fmt.Println("Success update of file")
	}

}

func Read_from_file(filename string) int {
	content, err := ioutil.ReadFile(filename)
	int_content, err := strconv.Atoi(string(content))
	if err != nil {
		fmt.Println("Unable to read file")
		log.Fatal(err)
	} else {
		fmt.Println("Success read from file ")
	}

	return int_content

}

func Clear_file(filename string) {
	err := os.Truncate(filename, 0)
	if err != nil {
		fmt.Println("Truncate error")
		log.Fatal(err)
	} else {
		fmt.Println("Success clear file")
	}

}
