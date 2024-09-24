package main

import (
	"fmt"
	"log"
	"os"

	"udemy.course.dev/v1/bank"
	"udemy.course.dev/v1/info"
	"udemy.course.dev/v1/os_managment"
	"udemy.course.dev/v1/struct_worker"
)

// Entrypoint

func main() {

	for {
		var choice int
		info.Main_menu_loop_info()
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)
		if choice > 0 {
			switch choice {
			case 1:
				os_managment.Install_packages()

			case 2:
				input_key := os_managment.Check_user_input()
				input_value := os_managment.Check_os_env(input_key)
				fmt.Println("Your value :", input_value)

			case 3:
				bank.Bank_crud_entrypoint()
			case 4:
				os.Exit(0)
			case 5:
				err := os.WriteFile("bank.txt", []byte("1000"), 0644)

				if err != nil {
					fmt.Println("Unable to write file")
					log.Fatal(err)
				}
			case 6:
				pointer_test()

			case 7:
				struct_worker.Struct_input()
			}

		} else {
			fmt.Println("Incorrect value")
		}

	}

}
